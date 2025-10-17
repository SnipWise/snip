#!/bin/bash

# Test script for similarities API

BASE_URL="http://localhost:3500"

echo "=========================================="
echo "Testing Similarities API"
echo "=========================================="
echo ""

# Test 2: Get similarities data
echo -e "\n2. Getting similarities data from /similarities..."
curl -X GET "${BASE_URL}/similarities" \
  -H "Content-Type: application/json" | jq '.'

echo -e "\n\n=========================================="
echo "Test completed!"
echo "=========================================="
