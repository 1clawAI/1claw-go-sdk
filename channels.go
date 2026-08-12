package oneclaw

import (
	"context"
	"fmt"
)

// ChannelsService handles messaging channel operations.
type ChannelsService struct {
	client *Client
}

// Channel represents an external messaging channel for an agent.
type Channel struct {
	ID                         string                 `json:"id"`
	OrgID                      string                 `json:"org_id,omitempty"`
	AgentID                    string                 `json:"agent_id"`
	ChannelType                string                 `json:"channel_type"`
	ChannelName                string                 `json:"channel_name,omitempty"`
	WebhookPath                string                 `json:"webhook_path,omitempty"`
	WebhookURL                 string                 `json:"webhook_url,omitempty"`
	WebhookSecret              *string                `json:"webhook_secret,omitempty"`
	IsActive                   bool                   `json:"is_active"`
	SlashCommandsEnabled       *bool                  `json:"slash_commands_enabled,omitempty"`
	VoiceTranscriptionEnabled  *bool                  `json:"voice_transcription_enabled,omitempty"`
	UnifiedConversationID      *string                `json:"unified_conversation_id,omitempty"`
	IsHomePlatform             *bool                  `json:"is_home_platform,omitempty"`
	SenderAllowlist            []string               `json:"sender_allowlist,omitempty"`
	AutoRespondEnabled         *bool                  `json:"auto_respond_enabled,omitempty"`
	Metadata                   map[string]interface{} `json:"metadata,omitempty"`
	Config                     map[string]interface{} `json:"config,omitempty"`
	CreatedAt                  string                 `json:"created_at"`
	UpdatedAt                  string                 `json:"updated_at"`
}

// ChannelList is the response from listing channels.
type ChannelList struct {
	Channels []Channel `json:"channels"`
}

// CreateChannelParams are parameters for creating a channel.
type CreateChannelParams struct {
	ChannelType                string                 `json:"channel_type"`
	ChannelName                string                 `json:"channel_name,omitempty"`
	Config                     map[string]interface{} `json:"config,omitempty"`
	SlashCommandsEnabled       *bool                  `json:"slash_commands_enabled,omitempty"`
	VoiceTranscriptionEnabled  *bool                  `json:"voice_transcription_enabled,omitempty"`
}

// UpdateChannelParams are parameters for updating a channel.
type UpdateChannelParams struct {
	ChannelName                *string                `json:"channel_name,omitempty"`
	IsActive                   *bool                  `json:"is_active,omitempty"`
	Config                     map[string]interface{} `json:"config,omitempty"`
	AutoRespondEnabled         *bool                  `json:"auto_respond_enabled,omitempty"`
	SenderAllowlist            []string               `json:"sender_allowlist,omitempty"`
	SlashCommandsEnabled       *bool                  `json:"slash_commands_enabled,omitempty"`
	VoiceTranscriptionEnabled  *bool                  `json:"voice_transcription_enabled,omitempty"`
}

// ChannelMessage represents a message in a channel.
type ChannelMessage struct {
	ID                   string                 `json:"id"`
	ChannelID            string                 `json:"channel_id"`
	Direction            string                 `json:"direction"`
	ExternalChatID       string                 `json:"external_chat_id,omitempty"`
	ExternalMessageID    *string                `json:"external_message_id,omitempty"`
	SenderName           *string                `json:"sender_name,omitempty"`
	Content              string                 `json:"content"`
	MediaURL             *string                `json:"media_url,omitempty"`
	IsVoiceMessage       *bool                  `json:"is_voice_message,omitempty"`
	VoiceFileID          *string                `json:"voice_file_id,omitempty"`
	VoiceDurationSecs    *float64               `json:"voice_duration_secs,omitempty"`
	TranscriptionStatus  *string                `json:"transcription_status,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt            string                 `json:"created_at"`
}

// ChannelMessageList is the response from listing channel messages.
type ChannelMessageList struct {
	Messages []ChannelMessage `json:"messages"`
}

// Create registers a new messaging channel for an agent (human-only).
func (s *ChannelsService) Create(ctx context.Context, agentID string, params CreateChannelParams) (*Channel, error) {
	var result Channel
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/channels", agentID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// List lists all messaging channels for an agent.
func (s *ChannelsService) List(ctx context.Context, agentID string) (*ChannelList, error) {
	var result ChannelList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/channels", agentID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific channel.
func (s *ChannelsService) Get(ctx context.Context, agentID, channelID string) (*Channel, error) {
	var result Channel
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/channels/%s", agentID, channelID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a channel (human-only).
func (s *ChannelsService) Update(ctx context.Context, agentID, channelID string, params UpdateChannelParams) (*Channel, error) {
	var result Channel
	err := s.client.doJSON(ctx, "PATCH", fmt.Sprintf("/v1/agents/%s/channels/%s", agentID, channelID), params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a channel (human-only).
func (s *ChannelsService) Delete(ctx context.Context, agentID, channelID string) error {
	return s.client.doJSON(ctx, "DELETE", fmt.Sprintf("/v1/agents/%s/channels/%s", agentID, channelID), nil, nil)
}

// SendMessage sends an outbound message via a registered channel.
func (s *ChannelsService) SendMessage(ctx context.Context, agentID, channelID string, content string, metadata map[string]interface{}) (*ChannelMessage, error) {
	body := map[string]interface{}{"content": content}
	if metadata != nil {
		body["metadata"] = metadata
	}
	var result ChannelMessage
	err := s.client.doJSON(ctx, "POST", fmt.Sprintf("/v1/agents/%s/channels/%s/send", agentID, channelID), body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListMessages lists inbound and outbound messages for a channel.
func (s *ChannelsService) ListMessages(ctx context.Context, agentID, channelID string) (*ChannelMessageList, error) {
	var result ChannelMessageList
	err := s.client.doJSON(ctx, "GET", fmt.Sprintf("/v1/agents/%s/channels/%s/messages", agentID, channelID), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
