'use server'
 
 
export async function authenticate(_currentState: unknown, formData: FormData) {
  try {
    1+2
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