# Go CLI Exercises — Dates & Time

These exercises are meant to **stretch your thinking**, not just your Go syntax. Treat them as small projects. Design first, code second.

---

## 1. 🗓️ Personal Date Utilities CLI

### Idea

Build a CLI tool that answers common **date math** questions.

The tool should feel like a small Unix utility: one binary, multiple commands.

### Example Usage

```bash
dateutil days-until 2026-03-15
dateutil age 1995-08-21
dateutil week 2026-01-03
dateutil diff 2026-01-01 2026-02-10
```

### Expected Behaviors

* `days-until`: Number of days from *today* until the given date
* `age`: Age in full years based on today
* `week`: ISO-8601 week number
* `diff`: Difference between two dates (days, maybe weeks later)

### Things to Think About

* What is "today"? Local time? UTC?
* What happens if the date is in the past?
* What should invalid input look like to the user?

### What You’ll Practice

* `time.Time`
* `time.Parse`
* `time.Duration`
* Formatting and parsing layouts (`2006-01-02`)
* Thinking about **time zones** and **leap years**

### Why This Is a Good Exercise

Go’s `time` package is famously strict. Fighting it a little now will make you much more confident later.

---

## 2. ⏱️ Time-Tracked Todo CLI (Upgrade Project)

### Idea

Extend your existing **todo CLI** with time tracking instead of starting from scratch.

Each task can optionally have time associated with it.

### Example Usage

```bash
todo start 3
todo stop 3
todo report --day
todo report --week
```

### Possible Features

* Start / stop timers for tasks
* Automatic timestamps
* Accumulated time per task
* Daily or weekly summaries

### Questions Worth Answering

* Can multiple tasks be running at once?
* What happens if you forget to stop a timer?
* Do you store raw timestamps or derived durations?

### What You’ll Practice

* Persisting timestamps (JSON or SQLite)
* Struct evolution over time
* Aggregating `time.Duration`
* Designing CLI subcommands

### Why This Is a Good Exercise

This feels like a *real tool*. It introduces real-world concerns like state, consistency, and reporting — without needing a web UI.

---

## Suggested Workflow

* Design the commands **before** writing code
* Start with a simple data model and let it evolve
* Prefer clarity over cleverness
* Use `go build` once the tool feels usable

---

## Reminder

These are not about finishing quickly.
They are about thinking clearly, naming things well, and building tools you’d actually want to use.
