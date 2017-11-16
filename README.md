# GoOST

[![Build Status](https://travis-ci.org/FX-HAO/GoOST.svg?branch=master)](https://travis-ci.org/FX-HAO/GoOST)

GoOST is an Order Statistic Tree implementation in golang.

An order statistic tree is a tree structure with two additional methods,
rank(node) and select(index), which allow array-like access in O(log n) time.
The idea is to simply store, for each node, the number of child-nodes.
Self-balancing search trees maintain a minimal height by rebalancing their
sub-trees.
