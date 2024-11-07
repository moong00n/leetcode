import assert from "node:assert";
import { describe, it } from "node:test";

const testCases = [
  { input: "leet**cod*e", expected: "lecoe" },
  { input: "erase*****", expected: "" },
  { input: "abc*def*ghi*", expected: "abdegh" },
];

function solution(s) {
  const stack = [];

  for (const char of s) {
    if (char !== "*") {
      stack.push(char);
    } else if (stack.length) {
      stack.pop();
    }
  }

  return stack.join("");
}

describe("Remove stars from string", () => {
  testCases.forEach(({ input, expected }) => {
    it(`"${input}" should return "${expected}"`, () => {
      assert.strictEqual(solution(input), expected);
    });
  });
});
