<script lang="ts">
  let { input }: { input: Input } = $props();

  const errorCss =
    "text-red-900 outline-red-300 placeholder:text-red-300 focus:outline-red-600";
  const nonErrorCss =
    "text-gray-900 outline-gray-300 placeholder:text-gray-400 ";

  let css = nonErrorCss;
  let errorMessage = $state("");
  let inputValue = $state(input.value || "");

  $effect(() => {
    if (errorMessage != "") {
      css = errorCss;
    } else {
      css = nonErrorCss;
    }
  });

  function onUpdate() {
    if (input.validator) {
      errorMessage = input.validator(inputValue);

      if (errorMessage != "") {
        return;
      }
    }

    if (input.onupdate) {
      input.onupdate(inputValue);
    }
  }
</script>

<div>
  {#if input.label}
    <label for="email" class="block text-sm/6 font-medium text-gray-900"
      >{input.label}</label
    >
  {/if}

  <div class="mt-2">
    <input
      oninput={onUpdate}
      type={input.type}
      bind:value={inputValue}
      class="block w-full rounded-md bg-white px-3 py-1.5 text-base outline outline-1 -outline-offset-1 focus:outline focus:outline-2 focus:outline-black focus:-outline-offset-2 sm:text-sm/6"
      placeholder={input.placeholder ? input.placeholder : ""}
    />
  </div>

  {#if errorMessage != ""}
    <p class="mt-2 text-sm text-red-600" id="email-error">{errorMessage}</p>
  {/if}
</div>
