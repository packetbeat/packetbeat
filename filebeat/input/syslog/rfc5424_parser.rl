// Code generated by ragel DO NOT EDIT.
package syslog

import "fmt"

%%{
    machine syslog_rfc5424;
    write data;
    variable p p;
    variable pe pe;
}%%

func Parse5424(data []byte, event *event) {
    var p, cs int
    pe := len(data)
    tok := 0
    eof := len(data)
    %%{

      include common "common.rl";

      include syslog_rfc5424 "syslog_rfc5424.rl";

      write init;
      write exec;
    }%%
}
