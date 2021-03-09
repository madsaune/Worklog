#track-work

A simple cmdline tool for tracking time spent on a particular task.

## Features

- [ ] Use CTRL+D or CTRL+C to stop tracking
- [ ] Write to sqlite database
- [ ] Export function
- [ ] Be able to add more metadata

## Datastructure

```golang
type Project struct {
  Title    string
  Metadata map[string]string
  Start    time.Time
  Stop     time.Time
}
```


## Example

```bash
track-work start "myProject" "more" "meta" "data"
Started tracking...

Press CTRL+D to stop.
```