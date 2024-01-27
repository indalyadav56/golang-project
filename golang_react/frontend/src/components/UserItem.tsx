const UserItem = ({ user }: any) => {
  return (
    <div>
      <h1>{user.email}</h1>
      <button>Delete</button>
    </div>
  )
}

export default UserItem
