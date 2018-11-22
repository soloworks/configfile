package configfile

// Helps output lines for a config file
// A line is set out as:
// Setting = Value // Comment
//
// Default Spacing:
// 19 for Setting
// 55 for value

import "bytes"

// ConfigFile is a structure holding a config file
type ConfigFile struct {
	content      bytes.Buffer
	widthSetting int
	widthValue   int
}

// New returns a ConfigFile object, instanced with defaults
func New() *ConfigFile {
	return &ConfigFile{
		widthSetting: 18,
		widthValue:   55,
	}
}

// AppendLineBlank builds a uniform line structure for config files
func (c *ConfigFile) AppendLineBlank() {
	// Terminate Line
	c.content.WriteString("\r\n")
}

// AppendLineSeperator builds a uniform line structure for config files
func (c *ConfigFile) AppendLineSeperator() {
	// Comment Character
	c.content.WriteString("#")
	// Pad out the setting
	for x := 0; x <= c.widthSetting+c.widthValue; x++ {
		c.content.WriteString("-")
	}
	// Terminate Line
	c.content.WriteString("\r\n")
}

// AppendLineComment builds a uniform line structure for config files
func (c *ConfigFile) AppendLineComment(comment string) {
	// Comment Character
	c.content.WriteString("# ")
	// Output the comment
	c.content.WriteString(comment)
	// Terminate Line
	c.content.WriteString("\r\n")
}

// AppendLineSetting builds a uniform line structure for config files
func (c *ConfigFile) AppendLineSetting(setting string, value string, comment string, writeEmpty bool) {

	if value != "" || writeEmpty {
		// Write the Setting
		c.content.WriteString(setting)
		// Pad out the setting
		for x := len(setting); x <= c.widthSetting; x++ {
			c.content.WriteString(" ")
		}
		// Write the Delimiter
		c.content.WriteString(" = ")
		// Write the Value
		c.content.WriteString(value)
		// If a comment is present
		if comment != "" {
			// Pad out the value
			for x := len(value); x <= c.widthValue; x++ {
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
