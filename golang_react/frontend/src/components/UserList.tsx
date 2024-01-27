import React from 'react'
import UserItem from './UserItem'

const UserList = ({ users }: any) => {
  return (
    <>
      {users?.map((user: any, index: any) => (
        <UserItem key={index} user={user} />
      ))}
    </>
  )
}

export default UserList
