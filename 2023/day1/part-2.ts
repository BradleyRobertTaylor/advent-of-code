import { getfirstAndLastDigit } from '../helpers/firstAndLastDigit';
import { getFirstAndLastSpelledNum } from '../helpers/firstAndLastSpelled';
import { parseInput } from '../helpers/parseInput';

const restoreCalibrationValues = (inputPath: string) => {
  const input = parseInput(inputPath);

  return input.reduce((sum, currLine) => {
    return sum + calculateLineCalibration(currLine);
  }, 0);
};

const calculateLineCalibration = (line: string) => {
  const digit = getfirstAndLastDigit(line);
  const spelled = getFirstAndLastSpelledNum(line);

  let first = '';
  let last = '';
  if ('first' in digit && 'first' in spelled) {
    first =
      digit.first.index < spelled.first.index
        ? digit.first.num
        : spelled.first.num;
    last =
      digit.last.index > spelled.last.index ? digit.last.num : spelled.last.num;
  } else if ('first' in spelled) {
    first = spelled.first.num;
    last = spelled.last.num;
  } else if ('first' in digit) {
    first = digit.first.num;
    last = digit.last.num;
  }

  return Number(`${first}${last}`);
};

console.log(restoreCalibrationValues('./day1-input.txt'));
