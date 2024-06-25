'use client';

import { authenticate } from '@/app/lib/actions';
import { useFormState, useFormStatus } from 'react-dom';
import { Card, CardHeader, CardBody, Image } from "@nextui-org/react";
import { Button } from "@nextui-org/react";
import React, { useState } from 'react';

interface User {
  first_name: string;
  last_name: string;
  email: string;
  avatar: string;
  phone: string;
  token_: string;
  refresh_token: string;
  created_at: string;
  updated_at: string;
  user_id: string;
}

export default function Page() {
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [user, setUser] = useState<User | null>(null);

  const dispatch = async (formData: FormData) => {
    const result = await authenticate(null, formData);
    if (typeof result === 'string') {
      setErrorMessage(result);
    } else {
      setUser(result as User);
      setErrorMessage(null);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen">
      <Card className="w-full max-w-[440px] h-[300px] flex flex-col justify-center">
        <CardHeader className="pb-0 pt-2 px-4 flex-col items-center flex gap-3">
          <h4 className="font-bold text-large">Order Pick</h4>
          <p>Let's Start Sorting!</p>
        </CardHeader>
        <CardBody className="pb-0 pt-2 px-4 flex-col items-center flex gap-3">
          <form onSubmit={(event) => {
            event.preventDefault();
            const formData = new FormData(event.currentTarget);
            dispatch(formData);
          }} className="w-full flex flex-col items-center">
            <input 
              type="email" 
              name="email" 
              placeholder="Email" 
              required 
              className="w-3/4 mb-4 p-2 border border-gray-300 rounded"
            />
            <input 
              type="password" 
              name="password" 
              placeholder="Password" 
              required 
              className="w-3/4 mb-4 p-2 border border-gray-300 rounded"
            />
            <div className="w-3/4 mb-4 text-center">
              {errorMessage && <p className="text-red-500">{errorMessage}</p>}
            </div>
            <LoginButton />
          </form>
          {user && (
            <div className="mt-4 text-center">
              <p>Welcome, {user.first_name} {user.last_name}!</p>
              <p>Email: {user.email}</p>
              <Image src={user.avatar} alt="User Avatar" className="w-16 h-16 rounded-full" />
            </div>
          )}
        </CardBody>
      </Card>
    </div>
  );
}

function LoginButton() {
  const { pending } = useFormStatus();

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    if (pending) {
      event.preventDefault();
    }
  };

  return (
    <Button color="default" variant="bordered" aria-disabled={pending} type="submit" onClick={handleClick}>
      Login
    </Button>
  );
}
