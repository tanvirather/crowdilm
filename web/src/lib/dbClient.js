class DbClient {
  constructor(dbName, version = 1, stores = []) {
    this.db = null;
    this.dbName = dbName;
    this.version = version;
    this.stores = stores;
  }

  async open() {
    return new Promise((resolve, reject) => {
      const request = indexedDB.open(this.dbName, this.version);
      request.onupgradeneeded = event => {
        this.db = event.target.result;
        this.stores.forEach(store => {
          if (!this.db.objectStoreNames.contains(store.name)) {
            const objectStore = this.db.createObjectStore(store.name, {
              keyPath: store.keyPath || 'id',
              autoIncrement: store.autoIncrement || false
            });
            if (store.indexes) {
              store.indexes.forEach(index =>
                objectStore.createIndex(index, index, { unique: false })
              );
            }
          }
        });
      };
      request.onsuccess = event => {
        this.db = event.target.result;
        resolve(this.db);
      };
      request.onerror = event => reject(`Error opening DB: ${event.target.errorCode}`);
    });
  }

  count = async (storeName) => this.#toPromise(this.#store(storeName).count());
  getAll = async (storeName) => this.#toPromise(this.#store(storeName).getAll());
  get = async (storeName, key) => this.#toPromise(this.#store(storeName).get(key));

  getByIndex = async (storeName, indexName, value) => {
    const store = this.#store(storeName, 'readwrite');
    const index = store.index(indexName);
    return this.#toPromise(index.getAll(value));
  };

  add = async (storeName, data) => this.#toPromise(this.#store(storeName, 'readwrite').add(data));
  addList = async (storeName, dataList) => {
    const store = this.#store(storeName, 'readwrite');
    const promises = dataList.map(item => this.#toPromise(store.add(item)));
    return Promise.all(promises);
  };
  update = async (storeName, data) => this.#toPromise(this.#store(storeName, 'readwrite').put(data));
  // delete = async (storeName, key) => this.#toPromise(this.#store(storeName, 'readwrite').delete(key));
  clearStore = async (storeName) => this.#toPromise(this.#store(storeName, 'readwrite').clear());

  #store = (storeName, mode = 'readonly') => this.db.transaction(storeName, mode).objectStore(storeName);
  #toPromise = async (request) => new Promise((resolve, reject) => {
    request.onsuccess = () => resolve(request.result);
    request.onerror = () => reject(request.error);
  });
}

export const dbClient = new DbClient('crowdilm', 1, [
  { name: 'sura' },
  { name: 'quran' },
  { name: 'quranLine', indexes: ['line', 'aya', 'hizb', 'juz', 'manzil', 'page', 'ruku', 'surah'] }
]);
