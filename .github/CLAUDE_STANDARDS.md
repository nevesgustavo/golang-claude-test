# Claude Code Review Standards

This document defines the coding standards and review criteria used by Claude AI for automated code reviews in this project.

## 🎯 Review Philosophy

Claude acts as a **senior Go engineer** with expertise in:
- Production-grade Go development
- Security best practices
- Performance optimization
- Test-driven development
- Clean architecture principles

## 📋 Mandatory Requirements

### 1. Security (CRITICAL)
- ✅ No hardcoded credentials, API keys, or secrets
- ✅ All user inputs must be validated and sanitized
- ✅ Use parameterized queries for database operations
- ✅ No sensitive data (PII) in logs or error messages
- ✅ Proper error handling that doesn't leak system information
- ✅ Check for common vulnerabilities: SQL injection, XSS, CSRF
- ✅ Verify proper authentication and authorization

### 2. Testing (HIGH PRIORITY)
- ✅ Unit tests required for all new functions
- ✅ Test coverage >= 80% for new code
- ✅ Table-driven tests for multiple scenarios
- ✅ Test both success and error cases
- ✅ Test edge conditions and boundary values
- ✅ Integration tests for critical paths
- ✅ Benchmarks for performance-critical code

### 3. Error Handling (HIGH PRIORITY)
- ✅ Never ignore errors (no `_ = err`)
- ✅ Wrap errors with context using `fmt.Errorf` with `%w`
- ✅ Return errors instead of panicking in library code
- ✅ Use custom error types for domain-specific errors
- ✅ Validate all error paths are tested

### 4. Documentation (REQUIRED)
- ✅ All exported functions/types must have godoc comments
- ✅ Comments should explain "why", not "what"
- ✅ Include usage examples for complex APIs
- ✅ Document edge cases and limitations
- ✅ Keep README updated for API changes

### 5. Code Quality (MEDIUM PRIORITY)
- ✅ Functions should be small and focused (max 50 lines)
- ✅ Use meaningful variable names (no single letters except loop counters)
- ✅ Avoid deep nesting (max 3-4 levels)
- ✅ No code duplication (DRY principle)
- ✅ Follow Single Responsibility Principle
- ✅ Use interfaces for abstraction and testability

## 🔍 Review Focus Areas

### Priority Levels

#### 🔴 CRITICAL (Must Fix Before Merge)
1. Security vulnerabilities
2. Data race conditions
3. Resource leaks (goroutines, connections, files)
4. Breaking API changes without migration path
5. Incorrect error handling leading to data loss

#### 🟡 HIGH (Should Fix Before Merge)
1. Missing or insufficient tests
2. Improper error handling
3. Concurrency safety issues
4. Performance regressions
5. Missing documentation for public APIs

#### 🟢 MEDIUM (Fix or Create Issue)
1. Code duplication
2. Suboptimal naming
3. Code complexity (cyclomatic complexity > 10)
4. Minor documentation gaps
5. Non-critical performance improvements

#### ⚪ LOW (Nice to Have)
1. Code formatting (should be automated)
2. Minor style inconsistencies
3. Optimization opportunities

## 🛡️ Go-Specific Best Practices

### Concurrency
```go
// ✅ GOOD: Proper goroutine management
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

g, ctx := errgroup.WithContext(ctx)
g.Go(func() error {
    return doWork(ctx)
})
if err := g.Wait(); err != nil {
    return fmt.Errorf("work failed: %w", err)
}

// ❌ BAD: Unmanaged goroutine
go func() {
    doWork() // No error handling, no cancellation
}()
```

### Error Handling
```go
// ✅ GOOD: Proper error wrapping
if err := db.Query(ctx, query); err != nil {
    return fmt.Errorf("failed to query users: %w", err)
}

// ❌ BAD: Lost error context
if err := db.Query(ctx, query); err != nil {
    return err
}
```

### Resource Management
```go
// ✅ GOOD: Defer cleanup
file, err := os.Open("data.txt")
if err != nil {
    return err
}
defer file.Close()

// ❌ BAD: Manual cleanup (easy to forget)
file, err := os.Open("data.txt")
if err != nil {
    return err
}
// ... lots of code ...
file.Close() // Might not be reached if there's an early return
```

### Testing
```go
// ✅ GOOD: Table-driven tests
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero", 0, 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

## 🚫 Blocked Patterns

Claude will flag these patterns:

| Pattern | Reason | Alternative |
|---------|--------|-------------|
| `fmt.Println` | No structured logging | Use `log/slog` or `zerolog` |
| `panic()` in libraries | Crashes caller | Return error |
| `TODO`, `FIXME`, `HACK` | Technical debt | Create issue, remove comment |
| Ignored errors `_ = err` | Silent failures | Handle or log errors |
| `time.Sleep` in tests | Flaky tests | Use channels or mocks |
| Magic numbers | Hard to understand | Use named constants |

## 📊 Code Metrics

Claude will evaluate:

- **Cyclomatic Complexity**: Max 10 per function
- **Function Length**: Max 50 lines
- **File Length**: Max 500 lines
- **Test Coverage**: Min 80% for new code
- **Comment Ratio**: 10-30% (not too little, not too much)

## 🎓 Learning Resources

When Claude suggests improvements, it will reference:

- [Effective Go](https://go.dev/doc/effective_go)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Proverbs](https://go-proverbs.github.io/)

## 🔄 Review Process

1. **Automated Checks** (CI/CD runs first)
   - Linting (golangci-lint)
   - Tests (go test)
   - Security scan (gosec)
   - Coverage (go test -cover)

2. **Claude Review** (triggered by @claude or /claude-review)
   - Analyzes code changes
   - Checks against standards
   - Provides actionable feedback
   - Suggests improvements with examples

3. **Human Review** (required for merge)
   - Team member approval
   - Address Claude's critical/high priority items
   - Discuss architectural decisions

## 💡 Tips for Better Reviews

### For Authors
- Keep PRs small (< 400 lines)
- Write descriptive commit messages
- Add context in PR description
- Self-review before requesting Claude review
- Run linters locally first

### For Reviewers
- Use `@claude` for specific questions
- Ask Claude to focus on specific areas
- Combine automated and human review
- Learn from Claude's suggestions

## 📝 Example Review Requests

```markdown
@claude Please review this PR focusing on:
1. Security implications of the new authentication flow
2. Potential race conditions in the cache implementation
3. Test coverage for error cases
```

```markdown
@claude This is a performance optimization. Can you:
1. Verify the logic is correct
2. Suggest any additional optimizations
3. Recommend benchmarks to add
```

## 🎯 Success Criteria

A PR is ready to merge when:

- ✅ All automated checks pass (linting, tests, security)
- ✅ Claude review has no CRITICAL or HIGH priority issues
- ✅ Test coverage >= 80%
- ✅ Documentation is complete
- ✅ At least 1 human approval
- ✅ All conversations resolved

---

**Remember**: Claude is a tool to help you write better code. Use it to learn, improve, and maintain high standards. The goal is not just to pass the review, but to understand why certain practices are important.
