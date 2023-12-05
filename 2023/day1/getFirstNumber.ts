export const getFirstNumber = (line: string) => {
  const numMap: Record<string, string> = {
    one: '1',
    two: '2',
    three: '3',
    four: '4',
    five: '5',
    six: '6',
    seven: '7',
    eight: '8',
    nine: '9',
  };

  const numbers = ['0', ...Object.keys(numMap), ...Object.values(numMap)];

  let earliestIdx = line.length;
  let firstNum = '';

  numbers.forEach((currNum) => {
    let currIdx = line.indexOf(currNum);

    if (currIdx !== -1 && earliestIdx > currIdx) {
      earliestIdx = currIdx;
      firstNum = currNum;
    }
  });

  return numMap[firstNum] || firstNum;
};
