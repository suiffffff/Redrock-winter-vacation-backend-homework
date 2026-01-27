




function switchPage(targetPageId) {
  //隐藏所有页面
  const allPages = document.querySelectorAll('.page-container');
  allPages.forEach(page => {
    page.classList.add('hidden');
  });
  //切换指定页面
  const targetPage = document.getElementById(targetPageId);
  if (targetPage) {
    targetPage.classList.remove('hidden');
  } else {
    console.error('找不到页面:', targetPageId);
  }
}

// switchPage()


function setupLogin() {
  document.querySelector('.login-btn').addEventListener('click', async () => {
    const loginForm = document.querySelector('.login-form')
    const loginObj = serialize(loginForm, { hash: true, empty: true })
    // console.log(loginForm)
    // console.log(loginObj)

    try {
      const response = await fetch('https://redrockwork.free.beeceptor.com/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer <access_token>'
        },
        body: JSON.stringify(loginObj)
      })

      if (!response.ok) {
        throw new Error('网络请求错误，状态码：' + response.status);
      }

      const result = await response.json();

      if (result.code === 0) {
        console.log('登录成功', result);
        // 执行跳转或保存 Token
      } else {
        // 业务逻辑错误 (如密码不对)
        const toastDom = document.querySelector('.my-toast')
        if (toastDom) {
          const toast = new bootstrap.Toast(toastDom)
          toast.show()
        }
        // 这里可以直接打印后端返回的错误信息
        console.log(result)
        console.error('登录失败:', result.msg || '未知错误');
      }

    } catch (error) {
      // 捕获所有错误 (网络错误或解析错误)
      console.error('Catch Error:', error);
      const toastDom = document.querySelector('.my-toast')
      if (toastDom) {
        const toast = new bootstrap.Toast(toastDom)
        toast.show();
      }
    }
  })
}
setupLogin()





