import { toast } from "sonner"

/**
 * 弹出成功通知
 * @param title 标题
 * @param content 内容（可选）
 */
export const notificationSuccess = (title: string, content?: string) => {
  toast.success(title, {
    description: content,
  })
}

/**
 * 弹出错误通知
 * @param title 标题
 * @param content 内容（可选）
 */
export const notificationError = (title: string, content?: string) => {
  toast.error(title, {
    description: content,
  })
}

/**
 * 弹出警告通知
 * @param title 标题
 * @param content 内容（可选）
 */
export const notificationWarn = (title: string, content?: string) => {
  toast.warning(title, {
    description: content,
  })
}

/**
 * 弹出信息通知
 * @param title 标题
 * @param content 内容（可选）
 */
export const notificationInfo = (title: string, content?: string) => {
  toast.info(title, {
    description: content,
  })
}
