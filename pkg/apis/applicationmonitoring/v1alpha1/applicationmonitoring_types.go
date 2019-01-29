package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApplicationMonitoringSpec defines the desired state of ApplicationMonitoring
type ApplicationMonitoringSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	AlertChannels []AlertChannel
}

type AlertChannel struct {
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	Type *string `json:"type,omitempty" yaml:"type,omitempty"`
	Severity int `json:"severity,omitempty" yaml:"severity,omitempty"`
	Settings AlertChannelSettings `json:"settings,omitempty" yaml:"settings,omitempty"`
}

type AlertChannelSettings struct {
	To *string `json:"to,omitempty" yaml:"to,omitempty"`
	SendResolved bool `json:"send_resolved,omitempty" yaml:"send_resolved,omitempty"`
	From *string `json:"from,omitempty" yaml:"from,omitempty"`
	SmtpSmarthost *string `json:"smtp_smarthost,omitempty" yaml:"smtp_smarthost,omitempty"`
	Hello *string `json:"hello,omitempty" yaml:"hello,omitempty"`
	AuthUsername *string `json:"auth_username,omitempty" yaml:"auth_username,omitempty"`
	AuthPassword *string `json:"auth_password,omitempty" yaml:"auth_password,omitempty"`
	AuthSecret *string `json:"auth_secret,omitempty" yaml:"auth_secret,omitempty"`
	AuthIdentity *string `json:"auth_identity,omitempty" yaml:"auth_identity,omitempty"`
	RequireTls *string `json:"require_tls,omitempty" yaml:"require_tls,omitempty"`
	TlsConfig AlertChannelTLSConfig `json:"tls_config,omitempty" yaml:"tls_config,omitempty"`
	Html *string `json:"html,omitempty" yaml:"html,omitempty"`
	Text *string `json:"text,omitempty" yaml:"text,omitempty"`
	Headers []string `json:"headers,omitempty" yaml:"headers,omitempty"`
	UserKey *string `json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Token *string `json:"token,omitempty" yaml:"token,omitempty"`
	Title *string`json:"title,omitempty" yaml:"title,omitempty"`
	Message *string `json:"message,omitempty" yaml:"message,omitempty"`
	Url *string `json:"url,omitempty" yaml:"url,omitempty"`
	Priority *string `json:"priority,omitempty" yaml:"priority,omitempty"`
	Retry *string `json:"retry,omitempty" yaml:"retry,omitempty"`
	Expire *string `json:"expire,omitempty" yaml:"expire,omitempty"`
	HttpConfig AlertChannelHttpConfig `json:"http_config,omitempty" yaml:"http_config,omitempty"`
	ServiceKey *string `json:"service_key,omitempty" yaml:"service_key,omitempty"`
	Client *string `json:"client,omitempty" yaml:"client,omitempty"`
	ClientUrl *string `json:"client_url,omitempty" yaml:"client_url,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Severity *string `json:"severity,omitempty" yaml:"severity,omitempty"`
	Details map[string]string `json:"details,omitempty" yaml:"details,omitempty"`
	Images []AlertChannelImage `json:"images,omitempty" yaml:"images,omitempty"`
	Links []AlertChannelLink `json:"links,omitempty" yaml:"links,omitempty"`
}

type AlertChannelHttpConfig struct {
	BasicAuth AlertChannelBasicAuth `json:"basic_auth,omitempty" yaml:"basic_auth,omitempty"`
	BearerToken *string `json:"bearer_token,omitempty" yaml:"bearer_token,omitempty"`
	BearerTokenFile *string `json:"bearer_token_file,omitempty" yaml:"bearer_token_file,omitempty"`
	TlsConfig AlertChannelTLSConfig `json:"tls_config,omitempty" yaml:"tls_config,omitempty"`
	ProxyUrl *string `json:"proxy_url,omitempty" yaml:"proxy_url,omitempty"`
}

type AlertChannelTLSConfig struct {
	CaFile *string `json:"ca_file,omitempty" yaml:"ca_file,omitempty"`
	CertFile *string `json:"cert_file,omitempty" yaml:"cert_file,omitempty"`
	KeyFile *string `json:"key_file,omitempty" yaml:"key_file,omitempty"` yaml:"metadata,omitempty"
	ServerName *string `json:"server_name,omitempty" yaml:"server_name,omitempty"`
	InsecureSkipVerify bool `json:"insecure_skip_verify,omitempty" yaml:"insecure_skip_verify,omitempty"`
}

type AlertChannelBasicAuth struct {
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `json:"password,omitempty" yaml:"password,omitempty"`
	PasswordFile *string  `json:"password_file,omitempty" yaml:"password_file,omitempty"`
}

type AlertChannelImage struct {
	Source *string `json:"source,omitempty" yaml:"source,omitempty"`
	Alt *string `json:"alt,omitempty" yaml:"alt,omitempty"`
	Text *string `json:"text,omitempty" yaml:"text,omitempty"`
}

type AlertChannelLink struct {
	Href *string `json:"href,omitempty" yaml:"href,omitempty"`
	Text *string `json:"text,omitempty" yaml:"text,omitempty"`
}

// ApplicationMonitoringStatus defines the observed state of ApplicationMonitoring
type ApplicationMonitoringStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApplicationMonitoring is the Schema for the applicationmonitorings API
// +k8s:openapi-gen=true
type ApplicationMonitoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationMonitoringSpec   `json:"spec,omitempty"`
	Status ApplicationMonitoringStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApplicationMonitoringList contains a list of ApplicationMonitoring
type ApplicationMonitoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationMonitoring `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApplicationMonitoring{}, &ApplicationMonitoringList{})
}
