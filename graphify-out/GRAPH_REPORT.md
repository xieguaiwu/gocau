# Graph Report - .  (2026-07-29)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 10 nodes · 18 edges · 2 communities (1 shown, 1 thin omitted)
- Extraction: 100% EXTRACTED · 0% INFERRED · 0% AMBIGUOUS
- Token cost: 120 input · 138 output

## Community Hubs (Navigation)
- Fitness Metrics Calculation
- Main Program Usage

## God Nodes (most connected - your core abstractions)
1. `handler()` - 9 edges
2. `Usage()` - 3 edges
3. `main()` - 3 edges
4. `readNumbers()` - 2 edges
5. `calAvg()` - 2 edges
6. `readWeightedNumbers()` - 2 edges
7. `calWAvg()` - 2 edges
8. `readFitnessMetric()` - 2 edges
9. `calFMetric()` - 2 edges

## Surprising Connections (you probably didn't know these)
- `handler()` --calls--> `Usage()`  [EXTRACTED]
  gocau.go → gocau.go  _Bridges community 1 → community 0_

## Import Cycles
- None detected.

## Communities (2 total, 1 thin omitted)

### Community 0 - "Fitness Metrics Calculation"
Cohesion: 0.46
Nodes (7): calAvg(), calFMetric(), calWAvg(), handler(), readFitnessMetric(), readNumbers(), readWeightedNumbers()

## Knowledge Gaps
- **1 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `handler()` connect `Fitness Metrics Calculation` to `Main Program Usage`?**
  _High betweenness centrality (0.375) - this node is a cross-community bridge._