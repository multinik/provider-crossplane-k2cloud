/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/multinik/provider-crossplane-k2cloud/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal crossplane-k2cloud credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		var creds map[string]any
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]any{
			"access_key": creds["access_key"].(string),
			"secret_key": creds["secret_key"].(string),
			"region":     creds["region"].(string),
			"endpoint":   creds["endpoint"].(string),
			"endpoints": []any{
				creds["endpoints"].(map[string]any),
			},
		}
		// Generate AWS_<SERVICE>_ENDPOINT env vars from endpoints map
		if eps, ok := creds["endpoints"].(map[string]any); ok {
			for svc, url := range eps {
				envVar := fmt.Sprintf("AWS_%s_ENDPOINT=%s", strings.ToUpper(svc), url.(string))
				ps.Environment = append(ps.Environment, envVar)
			}
		}
        // Propagate AWS credentials and default region as env vars
        ps.Environment = append(ps.Environment,
            fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", creds["access_key"].(string)),
            fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", creds["secret_key"].(string)),
            fmt.Sprintf("AWS_REGION=%s", creds["region"].(string)), // ensure 'region' key exists in creds
        )
		return ps, nil
	}
}
