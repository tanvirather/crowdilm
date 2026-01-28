export class StorageClient {
  constructor() {
    // Initialize with default settings if not present
    if (!this.getSetting()) {
      this.setSetting({
        selectedQuran1: "en.sahih",
        selectedQuranSize1: "24",
        selectedQuran2: "simple",
        selectedQuranSize2: "36",
        paging: "page",
      });
    }
    if (!this.getCurrentPage()) {
      this.setCurrentPage(1);
    }
  }

  getSetting() {
    return JSON.parse(localStorage.getItem("setting"));
  }
  setSetting(value) {
    localStorage.setItem("setting", JSON.stringify(value));
  }
  getCurrentPage() {
    return localStorage.getItem("currentPage");
  }
  setCurrentPage(value) {
    localStorage.setItem("currentPage", value);
  }
}

export const storageClient = new StorageClient();
