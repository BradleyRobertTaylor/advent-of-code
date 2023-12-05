import { getInput } from '../helpers/getInput';

const restoreCalibrationValues = (inputPath: string) => {
  const input = getInput(inputPath);

  return input.reduce((sum, currLine) => {
    return sum + calculateLineCalibration(currLine);
  }, 0);
};

const calculateLineCalibration = (line: string) => {
  let first = line.split('').find((char) => !isNaN(Number(char)))!;

  let last = '';
  for (let i = line.length - 1; i >= 0; i--) {
    if (!isNaN(Number(line[i]))) {
      last = line[i]!;
      break;
    }
  }

  return Number(`${first}${last}`);
};

console.log(restoreCalibrationValues('./day1-input.txt'));
