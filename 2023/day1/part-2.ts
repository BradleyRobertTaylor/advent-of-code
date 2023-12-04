import { getfirstAndLastDigit } from '../helpers/firstAndLastDigit';
import { getFirstAndLastSpelledNum } from '../helpers/firstAndLastSpelled';
import { parseInput } from '../helpers/parseInput';

const restoreCalibrationValues = async (inputPath: string) => {
  const input = await parseInput(inputPath);

  return input.reduce((sum, currLine) => {
    return sum + calculateLineCalibration(currLine);
  }, 0);
};

const calculateLineCalibration = (line: string) => {
  if (!line.trim()) return 0;

  const digit = getfirstAndLastDigit(line);
  const spelled = getFirstAndLastSpelledNum(line);

  if ('first' in digit && 'first' in spelled) {
    const first =
      digit.first.index < spelled.first.index
        ? digit.first.num
        : spelled.first.num;
    const last =
      digit.last.index > spelled.last.index ? digit.last.num : spelled.last.num;

    return Number(`${first}${last}`);
  } else if ('first' in spelled) {
    return Number(`${spelled.first.num}${spelled.last.num}`);
  } else if ('first' in digit) {
    return Number(`${digit.first.num}${digit.last.num}`);
  }

  return 0;
};

restoreCalibrationValues('./day1-input.txt').then(console.log);
