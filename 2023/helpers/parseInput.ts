import fs from 'node:fs/promises';

export const parseInput = async (inputPath: string) => {
  const input = await fs.readFile(inputPath, {
    encoding: 'utf8',
  });
  return input.split('\n');
};
