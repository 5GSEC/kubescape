package imagescan

import (
	"errors"
	"testing"
	"time"

	"github.com/anchore/grype/grype/db"
	"github.com/anchore/grype/grype/vulnerability"
	"github.com/stretchr/testify/assert"
)

// import (
// 	"context"
// 	"testing"

// 	"github.com/anchore/grype/grype/db"
// 	grypedb "github.com/anchore/grype/grype/db/v5"
// 	"github.com/anchore/grype/grype/match"
// 	"github.com/anchore/grype/grype/pkg"
// 	"github.com/anchore/grype/grype/presenter/models"
// 	"github.com/anchore/grype/grype/vulnerability"
// 	syftPkg "github.com/anchore/syft/syft/pkg"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// func TestNewScanService(t *testing.T) {
// 	dbCfg, _ := NewDefaultDBConfig()

// 	svc := NewScanService(dbCfg)

// 	assert.IsType(t, Service{}, svc)
// }

// func TestRegistryCredentials(t *testing.T) {
// 	tt := []struct {
// 		name     string
// 		username string
// 		password string
// 		want     bool
// 	}{
// 		{
// 			name:     "Valid credentials should not be empty",
// 			username: "user",
// 			password: "pass",
// 			want:     false,
// 		},
// 		{
// 			name:     "Empty username should be considered empty credentials",
// 			username: "",
// 			password: "pass",
// 			want:     true,
// 		},
// 		{
// 			name:     "Empty password should be considered empty credentials",
// 			username: "user",
// 			password: "",
// 			want:     true,
// 		},
// 		{
// 			name:     "Empty user and password should be considered empty credentials",
// 			username: "",
// 			password: "",
// 			want:     true,
// 		},
// 	}
// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			creds := RegistryCredentials{Username: tc.username, Password: tc.password}

// 			got := creds.IsEmpty()

// 			assert.Equal(t, tc.want, got)
// 		})
// 	}
// }

// func TestScan(t *testing.T) {
// 	tt := []struct {
// 		name  string
// 		image string
// 		creds RegistryCredentials
// 	}{
// 		{
// 			name:  "Valid image name produces a non-nil scan result",
// 			image: "nginx",
// 		},
// 		{
// 			name:  "Scanning a valid image with provided credentials should produce a non-nil scan result",
// 			image: "nginx",
// 			creds: RegistryCredentials{
// 				Username: "test",
// 				Password: "password",
// 			},
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctx := context.Background()
// 			dbCfg, _ := NewDefaultDBConfig()
// 			svc := NewScanService(dbCfg)
// 			creds := RegistryCredentials{}

// 			scanResults, err := svc.Scan(ctx, tc.image, creds)

// 			assert.NoError(t, err)
// 			assert.IsType(t, &models.PresenterConfig{}, scanResults)
// 		})
// 	}
// }

// // fakeMetaProvider is a test double that fakes an actual MetadataProvider
// type fakeMetaProvider struct {
// 	vulnerabilities map[string]map[string][]grypedb.Vulnerability
// 	metadata        map[string]map[string]*grypedb.VulnerabilityMetadata
// }

// func newFakeMetaProvider() *fakeMetaProvider {
// 	d := fakeMetaProvider{
// 		vulnerabilities: make(map[string]map[string][]grypedb.Vulnerability),
// 		metadata:        make(map[string]map[string]*grypedb.VulnerabilityMetadata),
// 	}
// 	d.fillWithData()
// 	return &d
// }

// func (d *fakeMetaProvider) GetAllVulnerabilityMetadata() (*[]grypedb.VulnerabilityMetadata, error) {
// 	return nil, nil
// }

// func (d *fakeMetaProvider) GetVulnerabilityMatchExclusion(id string) ([]grypedb.VulnerabilityMatchExclusion, error) {
// 	return nil, nil
// }

// func (d *fakeMetaProvider) GetVulnerabilityMetadata(id, namespace string) (*grypedb.VulnerabilityMetadata, error) {
// 	return d.metadata[id][namespace], nil
// }

// func (d *fakeMetaProvider) fillWithData() {
// 	d.metadata["CVE-2014-fake-1"] = map[string]*grypedb.VulnerabilityMetadata{
// 		"debian:distro:debian:8": {
// 			ID:        "CVE-2014-fake-1",
// 			Namespace: "debian:distro:debian:8",
// 			Severity:  "medium",
// 		},
// 	}

// 	d.vulnerabilities["debian:distro:debian:8"] = map[string][]grypedb.Vulnerability{
// 		"neutron": {
// 			{
// 				PackageName:       "neutron",
// 				Namespace:         "debian:distro:debian:8",
// 				VersionConstraint: "< 2014.1.3-6",
// 				ID:                "CVE-2014-fake-1",
// 				VersionFormat:     "deb",
// 			},
// 		},
// 	}
// }

// func TestExceedsSeverityThreshold(t *testing.T) {
// 	thePkg := pkg.Package{
// 		ID:      pkg.ID(uuid.NewString()),
// 		Name:    "the-package",
// 		Version: "v0.1",
// 		Type:    syftPkg.RpmPkg,
// 	}

// 	matches := match.NewMatches()
// 	matches.Add(match.Match{
// 		Vulnerability: vulnerability.Vulnerability{
// 			ID:        "CVE-2014-fake-1",
// 			Namespace: "debian:distro:debian:8",
// 		},
// 		Package: thePkg,
// 		Details: match.Details{
// 			{
// 				Type: match.ExactDirectMatch,
// 			},
// 		},
// 	})

// 	tt := []struct {
// 		name           string
// 		failOnSeverity string
// 		matches        match.Matches
// 		expectedResult bool
// 	}{
// 		{
// 			name:           "No severity set should pass",
// 			failOnSeverity: "",
// 			matches:        matches,
// 			expectedResult: false,
// 		},
// 		{
// 			name:           "Fail severity higher than vulnerability should not fail",
// 			failOnSeverity: "high",
// 			matches:        matches,
// 			expectedResult: false,
// 		},
// 		{
// 			name:           "Fail severity equal to vulnerability should fail",
// 			failOnSeverity: "medium",
// 			matches:        matches,
// 			expectedResult: true,
// 		},
// 		{
// 			name:           "Fail severity below found vuln should fail",
// 			failOnSeverity: "low",
// 			matches:        matches,
// 			expectedResult: true,
// 		},
// 	}

// 	metadataProvider := db.NewVulnerabilityMetadataProvider(newFakeMetaProvider())

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			scanResults := &models.PresenterConfig{
// 				Matches:          tc.matches,
// 				MetadataProvider: metadataProvider,
// 			}
// 			inputSeverity := vulnerability.ParseSeverity(tc.failOnSeverity)
// 			ours := ExceedsSeverityThreshold(scanResults, inputSeverity)

// 			assert.Equal(t, tc.expectedResult, ours)
// 		})
// 	}
// }

func TestParseSeverity(t *testing.T) {
	tests := []struct {
		name string
		want vulnerability.Severity
	}{
		{
			name: "",
			want: vulnerability.UnknownSeverity,
		},
		{
			name: "negligible",
			want: vulnerability.NegligibleSeverity,
		},
		{
			name: "low",
			want: vulnerability.LowSeverity,
		},
		{
			name: "medium",
			want: vulnerability.MediumSeverity,
		},
		{
			name: "high",
			want: vulnerability.HighSeverity,
		},
		{
			name: "critical",
			want: vulnerability.CriticalSeverity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseSeverity(tt.name)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		creds RegistryCredentials
		want  bool
	}{
		{
			name: "Both Non Empty",
			creds: RegistryCredentials{
				Username: "username",
				Password: "password",
			},
			want: false,
		},
		{
			name: "Password Empty",
			creds: RegistryCredentials{
				Username: "username",
				Password: "",
			},
			want: true,
		},
		{
			name: "Username Empty",
			creds: RegistryCredentials{
				Username: "",
				Password: "password",
			},
			want: true,
		},
		{
			name: "Both empty",
			creds: RegistryCredentials{
				Username: "",
				Password: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.creds.IsEmpty())
		})
	}
}

func TestNewDefaultDBConfig(t *testing.T) {
	config, shouldUpdate := NewDefaultDBConfig()
	assert.NotNil(t, config)
	assert.Equal(t, true, shouldUpdate)
	assert.Contains(t, config.DBRootDir, "grypedb")
	assert.Equal(t, "https://toolbox-data.anchore.io/grype/databases/listing.json", config.ListingURL)
}

func TestValidateDBLoad(t *testing.T) {
	currentTime := time.Now()
	tests := []struct {
		name               string
		loadErr            error
		status             *db.Status
		expectedErrMessage string
	}{
		{
			name:               "status nil",
			loadErr:            nil,
			status:             nil,
			expectedErrMessage: "unable to determine the status of the vulnerability db",
		},
		{
			name:    "loadErr nil and status error nil",
			loadErr: nil,
			status: &db.Status{
				Built:         currentTime,
				SchemaVersion: 7,
				Location:      "New Delhi",
				Checksum:      "invalid",
				Err:           nil,
			},
			expectedErrMessage: "",
		},
		{
			name:    "loadErr nil but status error not nil",
			loadErr: nil,
			status: &db.Status{
				Built:         currentTime,
				SchemaVersion: 7,
				Location:      "New Delhi",
				Checksum:      "invalid",
				Err:           errors.New("Some error"),
			},
			expectedErrMessage: "db could not be loaded: Some error",
		},
		{
			name:    "loadErr not nil",
			loadErr: errors.New("Some error"),
			status: &db.Status{
				Built:         currentTime,
				SchemaVersion: 7,
				Location:      "New Delhi",
				Checksum:      "invalid",
				Err:           errors.New("Some error"),
			},
			expectedErrMessage: "failed to load vulnerability db: Some error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDBLoad(tt.loadErr, tt.status)
			if err != nil {
				assert.Equal(t, tt.expectedErrMessage, err.Error())
			}
		})
	}
}

func TestGetProviderConfig(t *testing.T) {
	tests := []struct {
		name  string
		creds RegistryCredentials
	}{
		{
			name: "Both Non Empty",
			creds: RegistryCredentials{
				Username: "username",
				Password: "password",
			},
		},
		{
			name: "Password Empty",
			creds: RegistryCredentials{
				Username: "username",
				Password: "",
			},
		},
		{
			name: "Username Empty",
			creds: RegistryCredentials{
				Username: "",
				Password: "password",
			},
		},
		{
			name: "Both empty",
			creds: RegistryCredentials{
				Username: "",
				Password: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			providerConfig := getProviderConfig(tt.creds)
			assert.NotNil(t, providerConfig)
			assert.Equal(t, true, providerConfig.SynthesisConfig.GenerateMissingCPEs)
		})
	}
}

func TestNewScanService(t *testing.T) {
	defaultConfig, _ := NewDefaultDBConfig()
	svc := NewScanService(defaultConfig)
	assert.Equal(t, defaultConfig, svc.dbCfg)
}
