export interface LoginResponse {
    first_name:string;
    last_name:string;
    password:string;
    email:string;
    avatar:string;
    phone:string;
    token_:string;
    refresh_token:string;
    created_at:string;
    updated_at:string;
    user_id:string;
  }
  
  export async function login(email: string, password: string): Promise<LoginResponse> {
    const response = await fetch('http://localhost:8000/users/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });
  
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Something went wrong.');
    }
  
    return response.json();
  }