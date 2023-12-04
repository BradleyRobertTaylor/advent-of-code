import { CalibrationInfo } from './firstAndLastDigit';

const numMap = {
  one: '1',
  two: '2',
  three: '3',
  four: '4',
  five: '5',
  six: '6',
  seven: '7',
  eight: '8',
  nine: '9',
} as const;

export const getFirstAndLastSpelledNum = (
  line: string,
): { first: CalibrationInfo; last: CalibrationInfo } | {} => {
  if (Object.keys(numMap).every((num) => !line.includes(num))) return {};

  let first = Infinity;
  let firstNum = '';
  let last = -Infinity;
  let lastNum = '';

  Object.keys(numMap).forEach((num) => {
    const currFirst = line.indexOf(num);
    const currLast = line.lastIndexOf(num);

    if (currFirst < first && currFirst !== -1) {
      first = currFirst;
      firstNum = String(num);
    }

    if (currLast > last) {
      last = currLast;
      lastNum = String(num);
    }
  });

  return {
    first: { num: numMap[firstNum as keyof typeof numMap]!, index: first },
    last: { num: numMap[lastNum as keyof typeof numMap]!, index: last },
  };
};
