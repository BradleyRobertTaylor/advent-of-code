import fs from 'node:fs/promises';

export async function restoreCalibrationValues(inputPath: string) {
  const input = await fs.readFile(inputPath, {
    encoding: 'utf8',
  });
  const inputArr = input.split('\n');

  return inputArr.reduce((sum, currLine) => {
    return sum + calculateLineCalibration(currLine);
  }, 0);
}

function calculateLineCalibration(line: string): number {
  if (!line.trim()) return 0;

  let p1 = 0;
  let p2 = line.length - 1;

  while (isNaN(Number(line[p1]))) {
    p1++;
  }

  while (isNaN(Number(line[p2]))) {
    p2--;
  }

  return Number(`${line[p1]}${line[p2]}`);
}

interface LineCalibrationInfo {
  num: string | undefined;
  index: number;
}

function firstDigitIndex(line: string): LineCalibrationInfo | undefined {
  const chars = [...line];
  for (let i = 0; i < line.length; i++) {
    const char = chars[i];
    if (!isNaN(Number(char))) {
      return { num: char, index: i };
    }
  }
}

function lastDigitIndex(line: string): LineCalibrationInfo | undefined {
  const chars = [...line];
  for (let i = chars.length - 1; i >= 0; i--) {
    const char = chars[i];
    if (!isNaN(Number(char))) {
      return { num: char, index: i };
    }
  }
}

const numMap: Record<string, string> = {
  1: 'one',
  2: 'two',
  3: 'three',
  4: 'four',
  5: 'five',
  6: 'six',
  7: 'seven',
  8: 'eight',
  9: 'nine',
};

function findFirstSpelledOutNum(line: string) {
  let resultArr = new Array<number>(9).fill(0);
  for (const num in numMap) {
    const stringNum = numMap[num]!;
    const index = line.indexOf(stringNum);
    resultArr[Number(num) - 1] = index === -1 ? Infinity : index;
  }

  const max = Math.min(...resultArr);
}

function findLastSpelledOutNum(line: string) {}

findFirstSpelledOutNum('onetwo');
