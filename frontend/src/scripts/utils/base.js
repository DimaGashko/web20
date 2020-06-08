
export function readJsonData(type) {
   const jsonEl = document.querySelector(`.json-data[data-type="${type}"]`);
   if (!jsonEl) return null;

   return JSON.parse(jsonEl.innerHTML);
}