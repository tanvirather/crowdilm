export class SuraStore {
  #dbClient;
  #apiClient;
  #storeName = "sura";
  #url = "sura.php"

  constructor(dbClient, apiClient) {
    this.#dbClient = dbClient;
    this.#apiClient = apiClient
  }

  async getAll() {
    await this.#loadData();
    return await this.#dbClient.getAll(this.#storeName)
  }

  async #loadData() {
    var count = await this.#dbClient.count(this.#storeName);
    if (count === 0) {
      let dataList = await this.#apiClient.get(this.#url);
      await this.#dbClient.addList(this.#storeName, dataList);
    }
  }
}
