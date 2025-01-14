export class Mutex {
  private mutex = Promise.resolve();

  lock(): Promise<() => void> {
    return new Promise((resolve) => {
      this.mutex = this.mutex.then(() => new Promise(resolve));
    });
  }
}
