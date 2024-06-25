'use client'
 
import { authenticate } from '@/app/lib/actions'
import { useFormState, useFormStatus } from 'react-dom'
import {Card, CardHeader, CardBody, Image} from "@nextui-org/react";
import {Button} from "@nextui-org/react";

 
export default function Page() {
  const [errorMessage, dispatch] = useFormState(authenticate, undefined)
 
  return (

<div className="flex items-center justify-center min-h-screen">
  <Card className="w-full max-w-[440px] h-[600px] flex flex-col justify-center">
    <CardHeader className="pb-0 pt-2 px-4 flex-col items-center flex gap-3">
    <h4 className="font-bold text-large">Order Pick</h4>
    <p>Let's Start Sorting!</p>
    </CardHeader>
    <CardHeader className="pb-0 pt-2 px-4 flex-col items-center flex gap-3">
      <form action={dispatch} className="w-full flex flex-col items-center">
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
    </CardHeader>
  </Card>
</div>

  )
}
 
function LoginButton() {
  const { pending } = useFormStatus()
 
  const handleClick = (event) => {
    if (pending) {
      event.preventDefault()
    }
  }
 
  return (
    <Button color="default" variant="bordered"  aria-disabled={pending} type="submit" onClick={handleClick}>
      Login
  </Button>  )

}