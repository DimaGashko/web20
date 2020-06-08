import { getBaseHeaders, preprocessApiException } from '../utils/api';

/** 
 * @param {Post} post
 * @returns {Promise<Post>}
 */
export async function savePost(post) {
   return (post.slug) ? updatePost(post) : createPost(post);
}

/**
 * @param {Post} post
 * @returns {Promise<Post>}
 */
export async function createPost(post) {
   try {
      const resp = await fetch(`${process.env.API}/posts`, {
         method: 'POST',
         headers: getBaseHeaders(),
         body: JSON.stringify(post),
      });

      if (!resp.ok) {
         throw resp;
      }

      return await resp.json();
   } catch (e) {
      await preprocessApiException(e);
   }
}

/**
 * @param {Post} post
 * @returns {Promise<Post>}
 */
export async function updatePost(post) {
   try {
      const resp = await fetch(`${process.env.API}/posts/${post.slug}`, {
         method: 'PUT',
         headers: getBaseHeaders(),
         body: JSON.stringify(post),
      });

      if (!resp.ok) {
         throw resp;
      }

      return await resp.json();
   } catch (e) {
      await preprocessApiException(e);
   }
}