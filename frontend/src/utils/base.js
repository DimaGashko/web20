
/**
 * @param {HTMLElement} ell 
 * @param {string} className 
 * @param {bool} condition 
 */
export function setClassIf(ell, className, condition) {
   if (condition) {
      ell.classList.add(className);
   } else {
      ell.classList.remove(className);
   }
}