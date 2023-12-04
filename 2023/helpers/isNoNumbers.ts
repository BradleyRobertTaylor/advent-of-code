export const isNoNumbers = (str: string) => {
  return !/[0-9]+/g.test(str);
};
