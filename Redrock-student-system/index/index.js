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

switchPage('page-login')

// 登录页面事件设置
function setupLogin() {
  switchPage('page-login')
  document.querySelector('.login-btn').addEventListener('click', async () => {
    const loginForm = document.querySelector('.login-form')
    const loginObj = serialize(loginForm, { hash: true, empty: true })

    try {
      const response = await fetch('https://dd0e0bdc-b7fc-42f3-bc87-810ef0bd3eb3.mock.pstmn.io/login', {
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

        switchPage('page-system')
      } else {
        // 业务逻辑错误 (如密码不对)
        const toastDom = document.querySelector('.login-toast')
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
      const toastDom = document.querySelector('.login-toast')
      if (toastDom) {
        const toast = new bootstrap.Toast(toastDom)
        toast.show();
      }
    }
  })
  document.querySelector('.login-register-btn').addEventListener('click', () => {
    switchPage('page-register')
  })
}

// 注册页面事件设置
function setupRegister() {
  document.querySelector('.register-btn').addEventListener('click', async () => {
    const registerForm = document.querySelector('.register-form')
    const registerObj = serialize(registerForm, { hash: true, empty: true })

    try {
      console.log(registerObj)
      const { username, password, nickname, department } = registerObj

      if (!username || username.length < 8) {
        alert('账号长度必须大于8位')
        return
      }

      if (!password || password.length < 6) {
        alert('密码长度必须大于6位')
        return
      }

      if (!nickname || nickname.trim() === '') {
        alert('昵称不能为空')
        return
      }

      if (!department || department === '') {
        alert('请选择部门')
        return
      }

      // 检查账号是否重名
      try {
        const checkResponse = await fetch('https://dd0e0bdc-b7fc-42f3-bc87-810ef0bd3eb3.mock.pstmn.io/check-username', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ username })
        })

        if (!checkResponse.ok) {
          throw new Error('检查账号失败')
        }

        const checkResult = await checkResponse.json()
        console.log(checkResult.code)
        if (checkResult.code === 0) {
          // 账号已存在
          const toastDom = document.querySelector('.register-toast')
          if (toastDom) {
            // 修改提示信息
            const infoBox = toastDom.querySelector('.info-box')
            if (infoBox) {
              infoBox.textContent = '账号已存在'
            }
            const toast = new bootstrap.Toast(toastDom)
            toast.show()
          }
          return
        }

        // 账号不重名，发送注册请求
        console.log('账号可用，准备发送注册请求:', registerObj)

        const registerResponse = await fetch('https://dd0e0bdc-b7fc-42f3-bc87-810ef0bd3eb3.mock.pstmn.io/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(registerObj)
        })

        if (!registerResponse.ok) {
          throw new Error('注册请求失败，状态码：' + registerResponse.status)
        }

        const registerResult = await registerResponse.json()

        if (registerResult.code === 0) {
          console.log('注册成功:', registerResult)
          // 注册成功后可以跳转到登录页面
          alert('注册成功，请登录')
          switchPage('page-login')
        } else {
          console.log('注册失败:', registerResult)
          const toastDom = document.querySelector('.register-toast')
          if (toastDom) {
            const infoBox = toastDom.querySelector('.info-box')
            if (infoBox) {
              infoBox.textContent = registerResult.msg || '注册失败'
            }
            const toast = new bootstrap.Toast(toastDom)
            toast.show()
          }
        }

      } catch (checkError) {
        console.error('检查账号或注册时出错:', checkError)
        // 提示用户注册失败
        const toastDom = document.querySelector('.register-toast')
        if (toastDom) {
          const infoBox = toastDom.querySelector('.info-box')
          if (infoBox) {
            infoBox.textContent = '注册时出错，请稍后再试'
          }
          const toast = new bootstrap.Toast(toastDom)
          toast.show()
        }
      }

    }
    catch (error) {
      console.error('注册时出错:', error)
    }

  })
  document.querySelector('.comeback-btn').addEventListener('click', () => {
    switchPage('page-login')
  })
}

function setupSystem() {

}
// 页面加载时初始化所有事件监听器
function initApp() {
  setupLogin()
  setupRegister()
}

// 初始化应用
initApp()