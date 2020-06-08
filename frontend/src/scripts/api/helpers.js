import { preprocessApiException, getBaseHeaders } from '../utils/api';

/** 
 * @param {string} md
 * @returns {Promise<string>}
 */
export async function md2html(md) {
   try {
      const resp = await fetch(`${process.env.API}/md2html`, {
         method: 'POST',
         headers: getBaseHeaders(),
         body: JSON.stringify({ md }),
      });

      if (!resp.ok) {
         throw resp;
      }

      const { html } = await resp.json();
      return html;

   } catch (e) {
      await preprocessApiException(e);
   }
}