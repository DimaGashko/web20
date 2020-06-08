import { BASE_TXT } from '../texts';

export function getBaseHeaders() {
   return {
      'Content-Type': 'application/json',
   };
}

export async function preprocessApiException(exception) {
   throw (await getTextOfApiException(exception));
}

export async function getTextOfApiException(exception) {
   if (typeof exception === 'string') {
      return exception;
   }

   if ('ok' in exception) {
      console.log(exception);
      return `${exception.statusText}`;
   }

   console.error(exception);

   return BASE_TXT.NetworkError;
}