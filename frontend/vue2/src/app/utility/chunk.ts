function chunk(arr, chunkSize) {
  const res = [];
  while (arr.length > 0) {
    const c = arr.splice(0, chunkSize);
    // @ts-ignore
    res.push(c);
  }
  return res;
}

export {chunk}
