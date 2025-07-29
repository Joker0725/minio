import React from "react"
import { connect } from "react-redux"
import * as objectsActions from "./actions"

class ObjectTagsPanel extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      tags: {},             // { key: value }
      newKey: "",
      newValue: "",
      loading: true,
      error: null,
    }
  }

  componentDidMount() {
    this.loadTags()
  }

  loadTags = () => {
    const { bucket, object } = this.props
    this.setState({ loading: true })
    minioBrowser
      .sendJSONRequest("WebGetObjectTags", { bucket, object })
      .then((res) => {
        this.setState({ tags: res.tags || {}, loading: false })
      })
      .catch((err) => {
        this.setState({ error: "Failed to load tags", loading: false })
      })
  }

  handleAddTag = () => {
    const { tags, newKey, newValue } = this.state
    if (!newKey || !newValue) return
    const newTags = { ...tags, [newKey]: newValue }
    this.updateTags(newTags)
    this.setState({ newKey: "", newValue: "" })
  }

  handleDeleteTag = (key) => {
    const newTags = { ...this.state.tags }
    delete newTags[key]
    this.updateTags(newTags)
  }

  updateTags = (tags) => {
    const { bucket, object } = this.props
    minioBrowser
      .sendJSONRequest("WebPutObjectTags", {
        bucket,
        object,
        tags,
      })
      .then(() => {
        this.setState({ tags })
      })
      .catch(() => {
        this.setState({ error: "Failed to update tags" })
      })
  }

  render() {
    const { tags, newKey, newValue, loading, error } = this.state

    if (loading) return <div>Loading tags...</div>
    if (error) return <div className="alert alert-danger">{error}</div>

    return (
      <div className="tags-panel">
        <h5>Object Tags</h5>
        <table className="table table-striped">
          <thead>
            <tr>
              <th>Key</th>
              <th>Value</th>
              <th />
            </tr>
          </thead>
          <tbody>
            {Object.entries(tags).map(([key, value]) => (
              <tr key={key}>
                <td>{key}</td>
                <td>{value}</td>
                <td>
                  <button
                    className="btn btn-sm btn-danger"
                    onClick={() => this.handleDeleteTag(key)}
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
            <tr>
              <td>
                <input
                  type="text"
                  className="form-control"
                  placeholder="Key"
                  value={newKey}
                  onChange={(e) => this.setState({ newKey: e.target.value })}
                />
              </td>
              <td>
                <input
                  type="text"
                  className="form-control"
                  placeholder="Value"
                  value={newValue}
                  onChange={(e) => this.setState({ newValue: e.target.value })}
                />
              </td>
              <td>
                <button
                  className="btn btn-sm btn-success"
                  onClick={this.handleAddTag}
                >
                  Add
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    )
  }
}

const mapStateToProps = (state, ownProps) => ({
  bucket: state.objects.bucket,
  object: ownProps.object.name,
})

export default connect(mapStateToProps)(ObjectTagsPanel)
