export class QuranLineStore {
  #dbClient;
  #apiClient;
  #storeName = "quranLine";
  #url = "quranline.php"

  constructor(dbClient, apiClient) {
    this.#dbClient = dbClient;
    this.#apiClient = apiClient
  }

  async getByKey(selectedQuran1, selectedQuran2, index, key) {
    await this.#loadData(selectedQuran1, selectedQuran2);
    return await this.#dbClient.getByIndex(this.#storeName, index, `${key}`);
  }

  async #loadData(selectedQuran1, selectedQuran2) {
    var count = await this.#dbClient.count(this.#storeName);
    if (count === 0) {
      let dataList = await this.#apiClient.get(this.#url + `?quran_id=${selectedQuran1}`);
      let dataListWithId = dataList.map(obj => ({ id: `${selectedQuran1}_${obj.line}`, ...obj }));
      await this.#dbClient.addList(this.#storeName, dataListWithId);

      dataList = await this.#apiClient.get(this.#url + `?quran_id=${selectedQuran2}`);
      dataListWithId = dataList.map(obj => ({ id: `${selectedQuran2}_${obj.line}`, ...obj }));
      await this.#dbClient.addList(this.#storeName, dataListWithId);
    }
  }

}
