// Code generated by ragel DO NOT EDIT.
package syslog

%%{
    machine syslog;
    write data;
    variable p p;
    variable pe pe;
}%%

var (
  noDuplicates = []byte{'-', '.'}
)

// Parse parses Syslog events.
func Parse(data []byte, event *event) {
    var p, cs int
    pe := len(data)
    tok := 0
    eof := len(data)
    %%{

      include common "common.rl";
      include syslog_rfc3164 "syslog_rfc3164.rl";

      write init;
      write exec;
    }%%
}
