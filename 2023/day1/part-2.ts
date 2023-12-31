import { getInput } from '../helpers/getInput';
import { getFirstNumber } from './getFirstNumber';
import { getLastNumber } from './getLastNumber';

const restoreCalibrationValues = (inputPath: string) => {
  const input = getInput(inputPath);
  return input.reduce((sum, line) => sum + calculateLineCalibration(line), 0);
};

const calculateLineCalibration = (line: string) => {
  let first = getFirstNumber(line);
  let last = getLastNumber(line);

  return Number(`${first}${last}`);
};

console.log(restoreCalibrationValues('./day1-input.txt'));
