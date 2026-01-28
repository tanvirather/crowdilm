# dbClient

## Overview

`dbClient.md` documents the `dbClient` utility implemented in `dbClient.js`. `dbClient` is a thin IndexedDB wrapper providing simple async methods for common CRUD operations in the browser: `count`, `getAll`, `get`, `add`, `update`, and `delete`.

## Location

- **Source file**: `web/src/lib/dbClient.js`
- **Readme file**: `web/src/lib/dbClient.md` (this file)

## Requirements

- Runs in the browser environment (uses `indexedDB`). Not suitable for Node.js without an IndexedDB shim.

## API

- **Class**: `dbClient(dbName, version = 1)`

  - **Parameters**:
    - `dbName` (string): name of the IndexedDB database.
    - `version` (number, optional): DB version used for upgrades.

- **open(stores = [])**: Promise
  - Opens the database and creates object stores on `onupgradeneeded` per the `stores` array.
  - `stores` shape example:

```js
const stores = [
  {
    name: "lines",
    keyPath: "id", // optional (defaults to 'id')
    autoIncrement: true, // optional
    indexes: [{ name: "sura", keyPath: "sura", unique: false }],
  },
];
```

- **count(storeName)**: Promise<number>

  - Returns number of records in the store.

- **getAll(storeName)**: Promise<Array>

  - Returns all records from the store.

- **get(storeName, key = "id")**: Promise<any>

  - Gets a single record by key. Default key argument is the literal string `"id"` as used in the source; pass the actual key value when calling.

- **add(storeName, data)**: Promise<any>

  - Adds a record to the store (uses a readwrite transaction).

- **update(storeName, data)**: Promise<any>

  - Puts (adds or updates) a record by key into the store.

- **delete(storeName, key)**: Promise<any>
  - Deletes a record by key.

## Notes about internals

- The implementation wraps IDB requests into Promises with private helper methods (`#store` and `#toPromise`).
- Because the class uses private class fields and methods (e.g., `#store`, `#toPromise`), it requires a JS environment that supports those features (modern browsers or a transpiler).

## Usage example

```js
import dbClient from "./dbClient"; // adjust path as needed

const stores = [{ name: "lines", keyPath: "id", autoIncrement: true }];

const client = new dbClient("crowdilm-db", 1);

async function init() {
  await client.open(stores);

  // add
  const id = await client.add("lines", { text: "Example line", sura: 1 });

  // get
  const item = await client.get("lines", id);

  // getAll
  const all = await client.getAll("lines");

  // update
  await client.update("lines", { id, text: "Updated line" });

  // count
  const n = await client.count("lines");

  // delete
  await client.delete("lines", id);
}

init().catch(console.error);
```

## Export note

- The `dbClient.js` source shown in the repository defines the `dbClient` class but does not include an export statement in the snippet. If you intend to import it as an ES module, add the following at the end of `dbClient.js`:

```js
export default dbClient;
```

## Troubleshooting

- Error opening DB: Check browser console errors — common causes include invalid store definitions or insufficient permissions.
- If `request.onsuccess` never fires, ensure the browser supports IndexedDB and that other pages/tabs are not locking upgrades.
- If you get syntax errors about private fields (`#store`), run in a modern browser or transpile with Babel targeting a recent spec.

## Change log / Contributing

- Add issues or pull requests that document additional usage patterns or introduce export changes if needed.

## License

- Follow the repository's main license.
