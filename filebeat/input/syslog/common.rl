%%{
  machine common;
  action tok {
    fmt.Println(tok)
    fmt.Println(p)
    fmt.Println("---")
    tok = p
  }

  action priority {
    event.SetPriority(data[tok:p])
  }

  action message {
    event.SetMessage(data[tok:p])
  }

  action month {
    event.SetMonth(data[tok:p])
  }

  action year{
    event.SetYear(data[tok:p])
  }

  action month_numeric {
    event.SetMonthNumeric(data[tok:p])
  }

  action day {
    event.SetDay(data[tok:p])
  }

  action hour {
    event.SetHour(data[tok:p])
  }

  action minute {
    event.SetMinute(data[tok:p])
  }

  action second {
    event.SetSecond(data[tok:p])
  }

  action nanosecond{
    event.SetNanosecond(data[tok:p])
  }

  # NOTES: This allow to bail out of obvious non valid
  # hostname, this might not be ideal in all situation, but
  # when this happen we just go to the catch all case and at least
  # extract the message
  action lookahead_duplicates{
    if p-1 > 0 {
      for _, b := range noDuplicates {
        if data[p] == b && data[p-1] == b {
          p = tok -1
          fgoto catch_all;
        }
      }
    }
  }

  action hostname {
    event.SetHostname(data[tok:p])
  }

  action program {
    event.SetProgram(data[tok:p])
  }

  action pid {
    event.SetPid(data[tok:p])
  }

  action timezone {
    event.SetTimeZone(data[tok:p])
  }

  action sequence {
    event.SetSequence(data[tok:p])
  }

  action version{
    event.SetVersion(data[tok:p])
  }

  SP = ' ';

}%%
