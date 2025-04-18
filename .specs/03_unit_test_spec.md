
# ✅ Nuke - Unit Testing Specification

## 🎯 Goal

Ensure all core functions in the Nuke CLI app are covered with reliable and meaningful unit tests. This will help maintain app reliability, catch edge cases, and support safe future changes.

---

## 📁 Target Areas for Testing

### 1. 🔍 File Scanning
- **Function:** `ScanForRelatedFiles(appName string) ([]File, error)`
- **Tests should validate:**
  - Scans correct directories.
  - Finds files matching app name or bundle ID.
  - Returns correct number of files.
  - Handles apps not found gracefully.
  - Edge case: app name matches multiple apps.

---

### 2. 📦 Bundle Info Extraction
- **Function:** `ExtractBundleID(appPath string) (string, error)`
- **Tests should validate:**
  - Correctly parses `Info.plist`.
  - Handles missing or invalid `Info.plist`.
  - Returns accurate bundle ID.

---

### 3. 🗑 File Deletion
- **Function:** `DeleteFiles(files []File) error`
- **Tests should validate:**
  - All provided files are deleted.
  - Errors on permission issues are handled and reported.
  - Skips non-existent files without crashing.
  - Handles mixed permissions in file list.

---

### 4. 🧠 Filtering / User Exclusion Logic
- **Function:** `FilterFiles(files []File, exclusions []string) []File`
- **Tests should validate:**
  - Correctly removes excluded paths.
  - Is case-sensitive or insensitive as expected.
  - Handles empty or nil exclusions.

---

### 5. 📝 Logging & Reporting
- **Function:** `LogResults(results []DeleteResult)`
- **Tests should validate:**
  - All results are logged correctly.
  - Errors are included in output.
  - Summary is accurate (e.g., 5 deleted, 2 failed).

---

## 🧪 Testing Framework

Use Go's standard testing package:

```go
import "testing"
```

For assertions (optional but recommended):

```go
import "github.com/stretchr/testify/assert"
```

---

## 🔄 Coverage Expectations

- Aim for **80-90% test coverage**.
- Include:
  - Normal cases
  - Edge cases
  - Error paths
  - Empty inputs

Use coverage tools:

```bash
go test -cover ./...
```

---

## 📁 Suggested Folder Structure

```
nuke/
├── internal/
│   ├── scan/
│   │   └── scan_test.go
│   ├── bundle/
│   │   └── bundle_test.go
│   ├── delete/
│   │   └── delete_test.go
│   └── utils/
│       └── filter_test.go
```

---

## 📌 Summary

Unit tests are critical to Nuke's reliability. Every exported function should be tested with both valid and invalid input. Use table-driven tests where applicable for clean and scalable testing.

