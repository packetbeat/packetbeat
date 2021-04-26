// Code generated by ragel DO NOT EDIT.
package syslog

%%{
    machine format_check;
    write data;
    variable p p;
    variable pe pe;
}%%

func IsRFC5424Format(data []byte) bool {
    var p, cs int
    isRFC5424 := false
    pe := len(data)
    %%{

      action set_true {
          isRFC5424 = true
      }

      include common "common.rl";
      include syslog_rfc5424 "syslog_rfc5424.rl";

      main :=  ("<" PRIVAL_RANGE  ">" VERSION_RANGE SP digit{4}>set_true) ;

      write init;
      write exec;
    }%%
    return isRFC5424
}
