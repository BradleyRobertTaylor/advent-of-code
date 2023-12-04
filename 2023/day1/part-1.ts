import { isNoNumbers } from '../helpers/isNoNumbers';
import { parseInput } from '../helpers/parseInput';

const restoreCalibrationValues = async (inputPath: string) => {
  const input = await parseInput(inputPath);

  return input.reduce((sum, currLine) => {
    return sum + calculateLineCalibration(currLine);
  }, 0);
};

const calculateLineCalibration = (line: string) => {
  if (isNoNumbers(line)) return 0;

  let l = 0;
  let r = line.length - 1;

  while (isNaN(Number(line[l]))) {
    l++;
  }

  while (isNaN(Number(line[r]))) {
    r--;
  }

  return Number(`${line[l]}${line[r]}`);
};

restoreCalibrationValues('./day1-input.txt').then(console.log);
