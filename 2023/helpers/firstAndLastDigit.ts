export interface CalibrationInfo {
  num: string;
  index: number;
}

export const getfirstAndLastDigit = (
  line: string,
): { first: CalibrationInfo; last: CalibrationInfo } | {} => {
  if (!/[0-9]+/g.test(line)) return {};

  let p1 = 0;
  let p2 = line.length - 1;

  while (isNaN(Number(line[p1]))) {
    p1++;
  }

  while (isNaN(Number(line[p2]))) {
    p2--;
  }

  return {
    first: { num: line[p1]!, index: p1 },
    last: { num: line[p2]!, index: p2 },
  };
};
