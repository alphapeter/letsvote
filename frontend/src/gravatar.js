export const gravatar = {
  profilePicture (user, size) {
    if (!size) {
      size = 128
    }
    return 'https://www.gravatar.com/avatar/' + user.gravatar + '?s=' + size
  }
}
