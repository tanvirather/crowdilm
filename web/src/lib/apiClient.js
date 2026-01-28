class ApiClient {
  constructor(baseURL, defaultHeaders = {}) {
    this.baseURL = baseURL;
    this.defaultHeaders = { 'Content-Type': 'application/json', ...defaultHeaders };
  }

  get = async (endpoint, headers = {}) => this.#request(endpoint, 'GET', null, headers);
  post = async (endpoint, body, headers = {}) => this.#request(endpoint, 'POST', body, headers);
  put = async (endpoint, body, headers = {}) => this.#request(endpoint, 'PUT', body, headers);
  delete = async (endpoint, headers = {}) => this.#request(endpoint, 'DELETE', null, headers);

  async #request(endpoint, method = 'GET', body = null, headers = {}) {
    const url = `${this.baseURL}/${endpoint}`;
    const options = { method, headers: { ...this.defaultHeaders, ...headers } };
    if (body) { options.body = JSON.stringify(body); }

    try {
      const response = await fetch(url, options);
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(`Error ${response.status}: ${errorData.message || response.statusText}`);
      }
      return await response.json();
    } catch (error) {
      console.error('Fetch Error:', error);
      throw error;
    }
  }
}

export const apiClient = new ApiClient(import.meta.env.VITE_API_BASE_URL)
