<script>
  export let username
</script>

<span>{username.slice(0, 1).toUpperCase()}</span>

<style>
  span {
    width: 36px;
    height: 36px;
    border-radius: 50%;
    background-color: var(--secondary-violet);

    font-size: 14px;

    display: flex;
    justify-content: center;
    align-items: center;

    user-select: none;
  }
</style>