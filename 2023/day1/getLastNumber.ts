export const getLastNumber = (line: string) => {
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

  let latestIdx = -1;
  let lastNum = '';

  numbers.forEach((currNum) => {
    let currIdx = line.lastIndexOf(currNum);

    if (currIdx !== -1 && latestIdx < currIdx) {
      latestIdx = currIdx;
      lastNum = currNum;
    }
  });

  return numMap[lastNum] || lastNum;
};
