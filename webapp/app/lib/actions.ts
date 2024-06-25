'use server'
import { login, LoginResponse } from '../../usecase/use-login';
 
 
export async function authenticate(_currentState: unknown, formData: FormData): Promise<LoginResponse | string>  {
  try {
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;
    
    if (!email || !password) {
      throw new Error('CredentialsSignin');
    }
    
    const userData = await login(email, password);
    return userData;
  } catch (error) {
    if (error) {
      switch (error) {
        case 'CredentialsSignin':
          return 'Invalid credentials.'
        default:
          return 'Something went wrong.'
      }
    }
    throw error
  }
}