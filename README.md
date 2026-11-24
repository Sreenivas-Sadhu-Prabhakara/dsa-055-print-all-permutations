# Print All Permutations

> Striver SDE Sheet 055 · Day 10 Backtracking — in-place swap backtracking: position i tries every remaining candidate via a swap, then un-swaps. Go, stdlib only, tested.

## Problem

Generate every ordering of an array of distinct values — all n! of them.

## Approach

The swap formulation beats the used[]-array formulation on space: to fill position i, swap each element of the suffix into place, recurse on i+1, and swap back. The suffix *is* the "remaining candidates" set — no auxiliary bookkeeping, O(1) state per frame beyond the recursion itself. The symmetry of swap/un-swap guarantees the array returns to its original order after the full walk (and the copy appended at each leaf is essential — appending the live slice would alias every result to the same backing array). Trade-off: the output is not in lexicographic order; the factorial-system method (#054) or sorted-input insertion gives ordered output when that matters.

**Complexity:** O(n!·n) time (output size), O(n) extra space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
