# Claude Code Review Examples

This document provides real-world examples of how to use Claude Code effectively in different scenarios.

## 📚 Table of Contents

1. [Quick Start Examples](#quick-start-examples)
2. [Code Review Scenarios](#code-review-scenarios)
3. [Security Review Examples](#security-review-examples)
4. [Learning & Mentorship](#learning--mentorship)
5. [Troubleshooting Code](#troubleshooting-code)
6. [Best Practices](#best-practices)

## Quick Start Examples

### Example 1: Simple Function Review

**Scenario**: You added a new function and want quick feedback.

**Comment in PR**:
```markdown
@claude Can you review this multiply function?
```

**What Claude Will Check**:
- Function signature and naming
- Error handling (if applicable)
- Documentation
- Test coverage
- Edge cases

**Expected Response**:
```markdown
I reviewed the multiply function:

✅ **Verdict**: Needs Work

🟡 **High Priority**:
1. Missing unit tests - Add table-driven tests
2. No documentation - Add godoc comment

🟢 **Suggestions**:
- Consider adding overflow detection for large numbers

**Recommended Test**:
```go
func TestMultiply(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 6},
        {"negative", -2, 3, -6},
        {"zero", 0, 5, 0},
    }
    // ... test implementation
}
```
```

---

### Example 2: Full PR Review

**Scenario**: You have a complete feature ready for review.

**Comment in PR**:
```markdown
/claude-review
```

**What Claude Will Do**:
- Review all changed files
- Check security implications
- Verify test coverage
- Validate error handling
- Check documentation

---

### Example 3: Focused Security Review

**Scenario**: You're working with user authentication.

**Comment in PR**:
```markdown
@claude Please focus on security implications of this authentication change. 
Check for:
- SQL injection vulnerabilities
- Proper password hashing
- Session management
- Authorization checks
```

**What Claude Will Provide**:
- Security-specific analysis
- OWASP references
- Secure code examples
- Testing recommendations

---

## Code Review Scenarios

### Scenario 1: Adding a New API Endpoint

**Code**:
```go
func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    
    db.Exec("INSERT INTO users VALUES (?, ?)", user.Name, user.Email)
    
    w.WriteHeader(http.StatusCreated)
}
```

**Ask Claude**:
```markdown
@claude Review this new API endpoint for security and best practices
```

**Expected Feedback**:
```markdown
🔴 **Critical Issues**:

1. **SQL Injection Risk**: Using string interpolation in SQL query
   ```go
   // ❌ VULNERABLE
   db.Exec("INSERT INTO users VALUES (?, ?)", user.Name, user.Email)
   
   // ✅ SECURE - Use parameterized queries
   stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
   if err != nil {
       return fmt.Errorf("prepare statement: %w", err)
   }
   defer stmt.Close()
   ```

2. **No Input Validation**: User data not validated
   ```go
   // ✅ ADD VALIDATION
   if user.Email == "" || !isValidEmail(user.Email) {
       http.Error(w, "invalid email", http.StatusBadRequest)
       return
   }
   ```

3. **No Error Handling**: Decode and Exec errors ignored
   ```go
   // ✅ HANDLE ERRORS
   if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
       http.Error(w, "invalid request", http.StatusBadRequest)
       return
   }
   ```

🟡 **High Priority**:
- Add authentication middleware
- Implement rate limiting
- Add request logging
- Return proper JSON responses

🟢 **Suggestions**:
- Use a validation library (e.g., go-playground/validator)
- Add OpenAPI/Swagger documentation
- Implement request ID tracking
```

---

### Scenario 2: Refactoring Legacy Code

**Before**:
```go
func ProcessData(data string) string {
    result := ""
    for i := 0; i < len(data); i++ {
        if data[i] >= 'a' && data[i] <= 'z' {
            result = result + string(data[i]-32)
        } else {
            result = result + string(data[i])
        }
    }
    return result
}
```

**After**:
```go
func ProcessData(data string) string {
    return strings.ToUpper(data)
}
```

**Ask Claude**:
```markdown
@claude I refactored this function. Is the behavior preserved? Any issues?
```

**Expected Response**:
```markdown
✅ **Verdict**: LGTM with suggestions

**Analysis**:
The refactoring correctly preserves behavior - both convert lowercase to uppercase.

✅ **Improvements**:
1. Much simpler and more readable
2. Uses standard library (more maintainable)
3. Better performance (no string concatenation in loop)

🟢 **Suggestions**:

1. **Add documentation**:
   ```go
   // ProcessData converts all lowercase letters to uppercase.
   // Non-alphabetic characters are preserved.
   func ProcessData(data string) string {
       return strings.ToUpper(data)
   }
   ```

2. **Add tests to verify behavior**:
   ```go
   func TestProcessData(t *testing.T) {
       tests := []struct {
           input string
           want  string
       }{
           {"hello", "HELLO"},
           {"Hello123", "HELLO123"},
           {"", ""},
       }
       // ... test implementation
   }
   ```

3. **Consider edge cases**:
   - Unicode characters (strings.ToUpper handles these correctly)
   - Empty strings (handled)
   - Very large strings (consider streaming for huge inputs)
```

---

### Scenario 3: Adding Concurrency

**Code**:
```go
func ProcessItems(items []Item) error {
    var wg sync.WaitGroup
    
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            process(i)
        }(item)
    }
    
    wg.Wait()
    return nil
}
```

**Ask Claude**:
```markdown
@claude Review this concurrent implementation. Any race conditions or issues?
```

**Expected Feedback**:
```markdown
🟡 **High Priority Issues**:

1. **No Error Handling**: Errors from process() are lost
   ```go
   // ✅ BETTER: Collect errors
   func ProcessItems(items []Item) error {
       g, ctx := errgroup.WithContext(context.Background())
       
       for _, item := range items {
           item := item // Capture loop variable
           g.Go(func() error {
               return process(ctx, item)
           })
       }
       
       return g.Wait()
   }
   ```

2. **No Goroutine Limit**: Could spawn thousands of goroutines
   ```go
   // ✅ BETTER: Use worker pool
   func ProcessItems(items []Item) error {
       const maxWorkers = 10
       sem := make(chan struct{}, maxWorkers)
       
       // ... implementation with semaphore
   }
   ```

3. **No Context/Cancellation**: Can't stop processing
   ```go
   // ✅ ADD: Context support
   func ProcessItems(ctx context.Context, items []Item) error {
       // Check ctx.Done() in workers
   }
   ```

🟢 **Suggestions**:
- Add timeout for processing
- Add metrics/logging for monitoring
- Consider using a worker pool library
```

---

## Security Review Examples

### Example 1: Password Handling

**Code**:
```go
func CreateUser(username, password string) error {
    hash := md5.Sum([]byte(password))
    return db.Exec("INSERT INTO users VALUES (?, ?)", username, hash)
}
```

**Ask Claude**:
```markdown
@claude Security review please
```

**Expected Response**:
```markdown
🔴 **CRITICAL Security Issues**:

1. **Weak Hashing Algorithm**: MD5 is cryptographically broken
   ```go
   // ❌ INSECURE
   hash := md5.Sum([]byte(password))
   
   // ✅ SECURE: Use bcrypt
   import "golang.org/x/crypto/bcrypt"
   
   hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
   if err != nil {
       return fmt.Errorf("hash password: %w", err)
   }
   ```

2. **No Salt**: MD5 without salt is vulnerable to rainbow tables

3. **SQL Injection**: (See previous examples)

**Recommended Implementation**:
```go
func CreateUser(username, password string) error {
    // Validate input
    if len(password) < 8 {
        return errors.New("password too short")
    }
    
    // Hash with bcrypt (includes salt automatically)
    hash, err := bcrypt.GenerateFromPassword(
        []byte(password), 
        bcrypt.DefaultCost,
    )
    if err != nil {
        return fmt.Errorf("hash password: %w", err)
    }
    
    // Use parameterized query
    _, err = db.Exec(
        "INSERT INTO users (username, password_hash) VALUES (?, ?)",
        username,
        hash,
    )
    return err
}
```

**Testing**:
```go
func TestPasswordSecurity(t *testing.T) {
    password := "mySecretPassword123"
    
    // Hash twice with same password
    hash1, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    hash2, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    
    // Hashes should be different (salt is random)
    if string(hash1) == string(hash2) {
        t.Error("hashes should differ due to random salt")
    }
    
    // But both should verify correctly
    if err := bcrypt.CompareHashAndPassword(hash1, []byte(password)); err != nil {
        t.Error("hash1 should verify")
    }
}
```
```

---

## Learning & Mentorship

### Example 1: Understanding Code

**Scenario**: You inherited complex code and need explanation.

**Ask Claude**:
```markdown
@claude Can you explain how this caching mechanism works? 
I'm new to the codebase and want to understand the flow.

[paste code snippet]
```

**Expected Response**:
- Step-by-step explanation
- Diagram of data flow
- Potential issues to watch for
- Suggestions for improvements

---

### Example 2: Learning Best Practices

**Ask Claude**:
```markdown
@claude I'm learning Go. Can you review this code and teach me best practices?

func getData() ([]byte, error) {
    resp, _ := http.Get("https://api.example.com")
    body, _ := ioutil.ReadAll(resp.Body)
    return body, nil
}
```

**Expected Educational Response**:
```markdown
Great that you're learning Go! Let me help you improve this code:

🔴 **Critical Issues**:

1. **Ignored Errors**: Never ignore errors in Go
   ```go
   // ❌ BAD
   resp, _ := http.Get(url)
   
   // ✅ GOOD
   resp, err := http.Get(url)
   if err != nil {
       return nil, fmt.Errorf("fetch data: %w", err)
   }
   ```

2. **Resource Leak**: Response body not closed
   ```go
   // ✅ ALWAYS close response body
   defer resp.Body.Close()
   ```

3. **No Status Check**: Should verify HTTP status
   ```go
   if resp.StatusCode != http.StatusOK {
       return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
   }
   ```

**Improved Version**:
```go
func getData() ([]byte, error) {
    resp, err := http.Get("https://api.example.com")
    if err != nil {
        return nil, fmt.Errorf("http get: %w", err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
    }
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("read body: %w", err)
    }
    
    return body, nil
}
```

**Go Best Practices Applied**:
1. ✅ Always handle errors explicitly
2. ✅ Use defer for cleanup (resp.Body.Close)
3. ✅ Wrap errors with context (fmt.Errorf with %w)
4. ✅ Check HTTP status codes
5. ✅ Use io.ReadAll instead of deprecated ioutil.ReadAll

**Further Learning**:
- [Effective Go - Error Handling](https://go.dev/doc/effective_go#errors)
- [Go Blog - Error Handling](https://go.dev/blog/error-handling-and-go)
```

---

## Troubleshooting Code

### Example 1: Debugging a Bug

**Scenario**: Code has a subtle bug you can't find.

**Ask Claude**:
```markdown
@claude This function sometimes returns wrong results. Can you spot the bug?

func findDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    duplicates := []int{}
    
    for _, num := range nums {
        if seen[num] {
            duplicates = append(duplicates, num)
        }
        seen[num] = true
    }
    
    return duplicates
}

// Test case that fails:
// Input: [1, 2, 2, 2, 3]
// Expected: [2]
// Got: [2, 2]
```

**Expected Response**:
```markdown
Found the bug! 🐛

**Issue**: The function adds duplicates every time they're seen, not just once.

**Example**:
- Input: [1, 2, 2, 2, 3]
- First 2: Added to duplicates → [2]
- Second 2: Added again → [2, 2]
- Third 2: Added again → [2, 2, 2] (but you got [2, 2])

**Fix**:
```go
func findDuplicates(nums []int) []int {
    seen := make(map[int]bool)
    duplicates := make(map[int]bool) // Track which duplicates we've added
    result := []int{}
    
    for _, num := range nums {
        if seen[num] && !duplicates[num] {
            result = append(result, num)
            duplicates[num] = true
        }
        seen[num] = true
    }
    
    return result
}
```

**Better Approach** (more efficient):
```go
func findDuplicates(nums []int) []int {
    counts := make(map[int]int)
    
    // Count occurrences
    for _, num := range nums {
        counts[num]++
    }
    
    // Collect numbers that appear more than once
    duplicates := []int{}
    for num, count := range counts {
        if count > 1 {
            duplicates = append(duplicates, num)
        }
    }
    
    return duplicates
}
```

**Test**:
```go
func TestFindDuplicates(t *testing.T) {
    tests := []struct {
        input []int
        want  []int
    }{
        {[]int{1, 2, 2, 2, 3}, []int{2}},
        {[]int{1, 2, 3}, []int{}},
        {[]int{1, 1, 2, 2}, []int{1, 2}},
    }
    // ... test implementation
}
```
```

---

## Best Practices

### DO ✅

1. **Be Specific**
   ```markdown
   ✅ @claude Review error handling in the authentication flow
   ❌ @claude review this
   ```

2. **Provide Context**
   ```markdown
   ✅ @claude This is a performance-critical path handling 1000 req/sec. 
      Review for optimization opportunities.
   ❌ @claude is this fast?
   ```

3. **Ask for Learning**
   ```markdown
   ✅ @claude Explain why this approach is better and what I should learn
   ❌ @claude just fix it
   ```

4. **Focus Reviews**
   ```markdown
   ✅ @claude Focus on: 1) Security 2) Concurrency 3) Error handling
   ❌ @claude check everything
   ```

### DON'T ❌

1. **Don't Ask About Generated Code**
   ```markdown
   ❌ @claude review this auto-generated protobuf code
   ```

2. **Don't Review Huge PRs**
   ```markdown
   ❌ 2000 line PR with 50 files changed
   ✅ Break into smaller PRs (< 500 lines)
   ```

3. **Don't Ignore Critical Feedback**
   ```markdown
   ❌ Merging despite 🔴 CRITICAL security issues
   ✅ Address all critical issues before merge
   ```

4. **Don't Skip Automated Checks**
   ```markdown
   ❌ Asking Claude before running linters/tests
   ✅ Fix linter errors first, then ask Claude
   ```

---

## Quick Reference

### Common Commands

| Command | Use Case | Cost |
|---------|----------|------|
| `@claude review this` | Quick general review | ~$0.10 |
| `/claude-review` | Full PR review | ~$0.50 |
| `@claude explain [code]` | Understanding code | ~$0.05 |
| `@claude security review` | Security focus | ~$0.20 |
| `@claude suggest tests` | Test recommendations | ~$0.10 |

### Response Time

- Simple questions: ~10-30 seconds
- Full PR review: ~30-60 seconds
- Complex analysis: ~60-120 seconds

### When to Use Each Workflow

| Workflow | When to Use |
|----------|-------------|
| `claude.yaml` | Quick questions, specific reviews, learning |
| `claude-code-review.yaml` | Full PR reviews, pre-merge validation |
| `security-review.yaml` | Security-sensitive changes, auth code |

---

**Pro Tip**: Start with small, focused requests to learn how Claude responds, then gradually use it for larger reviews as you get comfortable with the feedback style.
