package configfile

// Helps output lines for a config file
// A line is set out as:
// Setting = Value // Comment
//
// Spacing:
// 24 for Setting
// 1 for =
// 49 for value
// 1 for space
// Rest for comment after //

import "bytes"

// ConfigFile is a structure holding a config file
type ConfigFile struct {
	content      bytes.Buffer
	widthSetting int
	widthValue   int
}

// NewConfigFile returns a ConfigFile object, instanced with defaults
func NewConfigFile() *ConfigFile {
	return &ConfigFile{
		widthSetting: 19,
		widthValue:   55,
	}
}

// AppendBlankLine builds a uniform line structure for config files
func (c *ConfigFile) AppendBlankLine() {
	// Terminate Line
	c.content.WriteString("\r\n")
}

// AppendSeperatorLine builds a uniform line structure for config files
func (c *ConfigFile) AppendSeperatorLine() {
	// Comment Character
	c.content.WriteString("#")
	// Pad out the setting
	for c.content.Len() < c.widthSetting+c.widthValue {
		c.content.WriteString("-")
	}
	// Terminate Line
	c.content.WriteString("\r\n")
}

// AppendCommentLine builds a uniform line structure for config files
func (c *ConfigFile) AppendCommentLine(comment string) {
	// Comment Character
	c.content.WriteString("# ")
	// Output the comment
	c.content.WriteString(comment)
	// Terminate Line
	c.content.WriteString("\r\n")
}

// AppendLine builds a uniform line structure for config files
func (c *ConfigFile) AppendLine(setting string, value string, comment string, writeEmpty bool) {

	if value != "" || writeEmpty {
		// Write the Setting
		c.content.WriteString(setting)
		// Pad out the setting
		for c.content.Len() < 19 {
			c.content.WriteString(" ")
		}
		// Write the Delimiter
		c.content.WriteString(" = ")
		// Write the Value
		c.content.WriteString(value)
		// If a comment is present
		if comment != "" {
			// Pad out the value
			for c.content.Len() < 55 {
				c.content.WriteString(" ")
			}
			// Write the Delimiter
			c.content.WriteString(" // ")
			// Write the Comment
			c.content.WriteString(comment)

		}
		// Terminate Line
		c.content.WriteString("\r\n")
	}
}

// String exports the current config file as a string
func (c *ConfigFile) String() string {
	return c.content.String()
}

// Bytes exports the current config file as a string
func (c *ConfigFile) Bytes() []byte {
	return c.content.Bytes()
}
