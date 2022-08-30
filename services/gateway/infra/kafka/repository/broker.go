package repository

func (c *Client) DeleteTopic(topic string) error {
	err := c.clusterAdmin.DeleteTopic(topic)
	if err != nil {
		return err
	}
	return nil
}
